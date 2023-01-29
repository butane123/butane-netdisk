package logic

import (
	"bytes"
	"cloud-disk/common/errorx"
	"cloud-disk/common/utils"
	"cloud-disk/service/repository/model"
	"cloud-disk/service/user/rpc/user"
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/tencentyun/cos-go-sdk-v5"

	"cloud-disk/service/repository/api/internal/svc"
	"cloud-disk/service/repository/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadByChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadByChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadByChunkLogic {
	return &FileUploadByChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 根据md5码查看文件是否存在，若存在则秒传成功。若不存在则插入数据库数据并拼接name和ext，上传文件到cos
func (l *FileUploadByChunkLogic) FileUploadByChunk(req *types.FileUploadByChunkRequest, file multipart.File, fileHeader *multipart.FileHeader) (resp *types.FileUploadByChunkResponse, err error) {
	// 判断是否已达用户容量上限
	userIdentity := json.Number(fmt.Sprintf("%v", l.ctx.Value("userIdentity"))).String()
	volumeInfo, err := l.svcCtx.UserRpc.FindVolumeByIdentity(l.ctx, &user.FindVolumeReq{Identity: userIdentity})
	if err != nil {
		return nil, err
	}
	if volumeInfo.NowVolume+fileHeader.Size > volumeInfo.TotalVolume {
		return nil, errorx.NewDefaultError("文件过大！")
	}
	_, err = l.svcCtx.UserRpc.AddVolume(l.ctx, &user.AddVolumeReq{
		Identity: userIdentity,
		Size:     fileHeader.Size,
	})
	if err != nil {
		return nil, err
	}
	newIdentity := utils.GenerateUUID()
	// 判断文件是否已存在，若已存在则为秒传成功
	b := make([]byte, fileHeader.Size)
	_, err = file.Read(b)
	if err != nil {
		return nil, err
	}
	md5Str := fmt.Sprintf("%x", md5.Sum(b))
	count, err := l.svcCtx.RepositoryPoolModel.CountByHash(l.ctx, md5Str)
	if count > 0 {
		repositoryInfo, err := l.svcCtx.RepositoryPoolModel.FindRepositoryPoolByHash(l.ctx, md5Str)
		if err != nil {
			return nil, err
		}
		return &types.FileUploadByChunkResponse{Identity: repositoryInfo.Identity.String}, err
	}
	// 上传文件到cos，并得到filepath
	// //先将文件分块
	err = GenerateChunk(file, fileHeader, md5Str)
	if err != nil {
		return nil, err
	}
	// //将分块后的文件进行上传
	name, baseName, err := UploadByPart(file, fileHeader, md5Str, newIdentity)
	if err != nil {
		return nil, err
	}
	filePath := utils.CosUrl + "/" + name
	// 插入数据
	_, err = l.svcCtx.RepositoryPoolModel.Insert(l.ctx, &model.RepositoryPool{
		Identity: sql.NullString{String: newIdentity, Valid: true},
		Hash:     sql.NullString{String: md5Str, Valid: true},
		Ext:      sql.NullString{String: path.Ext(baseName), Valid: true},
		Size:     sql.NullInt64{Int64: fileHeader.Size, Valid: true},
		Path:     sql.NullString{String: filePath, Valid: true},
	})
	if err != nil {
		return nil, errorx.NewDefaultError("上传失败！")
	}
	return &types.FileUploadByChunkResponse{Identity: newIdentity}, err
}

func GenerateChunk(file multipart.File, fileHeader *multipart.FileHeader, md5Str string) error {
	ChunkSize := utils.ChunkSize
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
func UploadByPart(file multipart.File, fileHeader *multipart.FileHeader, md5Str string, newIdentity string) (string, string, error) {
	//获得上传的Upload，表示现在将上传的文件
	ChunkSize := utils.ChunkSize
	chunkNum := math.Ceil(float64(fileHeader.Size) / float64(ChunkSize))
	u, _ := url.Parse(utils.CosUrl)
	bs := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(bs, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  utils.SecretID,
			SecretKey: utils.SecretKey,
		},
	})
	baseName := path.Base(fileHeader.Filename)
	name := "cloud-disk/" + newIdentity + baseName
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
	return name, baseName, nil
}
