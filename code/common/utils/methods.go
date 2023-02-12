package utils

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"golang.org/x/crypto/scrypt"

	"github.com/tencentyun/cos-go-sdk-v5"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

func GenerateVerificationCode() string {
	str, verificationCode := "0123456789", ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < VerificationCodeLength; i++ {
		verificationCode += fmt.Sprintf("%c", str[rand.Intn(10)])
	}
	return verificationCode
}

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func GenerateJwtToken(secreKey string, iat, seconds int64, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secreKey))
}

// 文件分块
func GenerateChunk(file multipart.File, fileHeader *multipart.FileHeader, md5Str string) error {
	ChunkSize := ChunkSize
	chunkNum := math.Ceil(float64(fileHeader.Size) / float64(ChunkSize))
	for i := 0; i < int(chunkNum); i++ {
		//新建块，初始化大小
		nowBlo := make([]byte, ChunkSize)
		file.Seek(int64(i*ChunkSize), 0)
		if int64(ChunkSize) > fileHeader.Size-int64(i*ChunkSize) {
			nowBlo = make([]byte, int64(ChunkSize)-(fileHeader.Size-int64(i*ChunkSize)))
		}
		//读入块数据，向nowBlow中读入file的数据
		file.Read(nowBlo)
		f, err := os.OpenFile("service/repository/filePath/"+md5Str+strconv.FormatInt(int64(i), 10)+".chunk", os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		//输出块
		f.Write(nowBlo)
		f.Close()
	}
	file.Close()
	return nil
}

// 文件上传
func CosUpload(fileHeader *multipart.FileHeader, newId int64, b []byte) (string, string, error) {
	u, _ := url.Parse(CosUrl)
	bs := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(bs, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SecretID,
			SecretKey: SecretKey,
		},
	})
	baseName := path.Base(fileHeader.Filename)
	name := "butane-netdisk/" + strconv.FormatInt(newId, 10) + baseName
	_, err := c.Object.Put(context.Background(), name, bytes.NewReader(b), nil)
	if err != nil {
		return "", "", err
	}
	filePath := CosUrl + "/" + name
	return filePath, baseName, nil
}

// 文件分块上传
func CosUploadByPart(fileHeader *multipart.FileHeader, md5Str string, newId int64) (string, string, error) {
	//获得上传的Upload，表示现在将上传的文件
	ChunkSize := ChunkSize
	chunkNum := math.Ceil(float64(fileHeader.Size) / float64(ChunkSize))
	u, _ := url.Parse(CosUrl)
	bs := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(bs, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SecretID,
			SecretKey: SecretKey,
		},
	})
	baseName := path.Base(fileHeader.Filename)
	name := "butane-netdisk/" + strconv.FormatInt(newId, 10) + baseName
	v, _, err := c.Object.InitiateMultipartUpload(context.Background(), name, nil)
	if err != nil {
		return "", "", err
	}
	UploadID := v.UploadID
	opt := &cos.CompleteMultipartUploadOptions{}
	for i := 0; i < int(chunkNum); i++ {
		//获得该块的md5码值PartETag
		f, err := os.ReadFile("service/repository/filePath/" + md5Str + strconv.FormatInt(int64(i), 10) + ".chunk")
		if err != nil {
			return "", "", err
		}
		resp, err := c.Object.UploadPart(
			context.Background(), name, UploadID, i+1, bytes.NewReader(f), nil,
		)
		if err != nil {
			return "", "", err
		}
		PartETag := resp.Header.Get("ETag")
		//将该块塞入块数组opt
		opt.Parts = append(opt.Parts, cos.Object{PartNumber: i + 1, ETag: PartETag})
	}
	//将所有块上传
	_, _, err = c.Object.CompleteMultipartUpload(
		context.Background(), name, UploadID, opt,
	)
	if err != nil {
		return "", "", err
	}
	filePath := CosUrl + "/" + name
	return filePath, baseName, nil
}

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}

func GenerateNewId(redis *redis.Redis, keyPrefix string) int64 {
	//获取当前时间戳
	nowStamp := time.Now().Unix() - BeginTimeStamp
	//调用lua脚本，获取当天累计序列号
	nowDate := time.Now().Format("2006:01:02")
	newKeyString := "cache:icr:" + keyPrefix + ":" + nowDate
	count, err := redis.Incr(newKeyString)
	if err != nil {
		fmt.Println("生成id错误！")
		return 0
	}
	//拼接结果
	return nowStamp<<IdCountBit | count
}
