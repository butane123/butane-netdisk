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
	"mime/multipart"
	"net/http"
	"net/url"
	"path"

	"github.com/tencentyun/cos-go-sdk-v5"

	"cloud-disk/service/repository/api/internal/svc"
	"cloud-disk/service/repository/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 根据md5码查看文件是否存在，若存在则秒传成功。若不存在则插入数据库数据并拼接name和ext，上传文件到cos
func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest, file multipart.File, fileHeader *multipart.FileHeader) (resp *types.FileUploadResponse, err error) {
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
		return &types.FileUploadResponse{Identity: repositoryInfo.Identity.String}, err
	}
	// 上传文件到cos，并得到filepath

	u, _ := url.Parse(utils.CosUrl)
	bs := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(bs, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  utils.SecretID,
			SecretKey: utils.SecretKey,
		},
	})
	newIdentity := utils.GenerateUUID()
	baseName := path.Base(fileHeader.Filename)
	name := "cloud-disk/" + newIdentity + baseName
	_, err = c.Object.Put(context.Background(), name, bytes.NewReader(b), nil)
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
	return &types.FileUploadResponse{Identity: newIdentity}, err
}
