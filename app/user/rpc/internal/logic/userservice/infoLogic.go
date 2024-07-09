package userservicelogic

import (
	"blog/common/helper"
	"blog/database/model"
	"context"

	"blog/app/user/rpc/internal/svc"
	"blog/app/user/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *rpc.InfoReq) (*rpc.User, error) {
	info := model.User{}
	err := l.svcCtx.DB.Where("id=?", in.UserId).First(&info).Error
	if err != nil {
		return nil, err
	}
	cInfo := rpc.User{}
	helper.ExchangeStruct(info, &cInfo)
	return &cInfo, nil
}
