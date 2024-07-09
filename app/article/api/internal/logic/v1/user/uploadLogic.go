package user

import (
	"blog/app/article/api/internal/svc"
	"blog/app/article/api/internal/types"
	"blog/common/response/errx"
	"blog/common/tencent/cos"
	"context"
	"io"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(r *http.Request) (resp *types.UploadRes, err error) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		return nil, errx.NewMessageError(err.Error())
	}
	defer file.Close()
	//创建上传目录
	os.Mkdir("./upload", os.ModePerm)
	//创建上传文件
	filePath := "./upload/" + handler.Filename
	f, err := os.Create(filePath)
	defer f.Close()
	io.Copy(f, file)

	client, err := cos.NewClient(l.svcCtx.Config.TencentCloud)
	url, err := client.UploadLocal(handler.Filename, filePath, "article")
	url, _ = client.GetPresignedURL(url)
	if err != nil {
		return nil, errx.NewMessageError(err.Error())
	}
	defer file.Close()
	return &types.UploadRes{File: url}, nil
}
