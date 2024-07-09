package menu

import (
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/database/model"
	"context"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.MenuInfoReq) (resp *types.Menu, err error) {
	info := model.AdminMenu{}
	err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.Id).First(&info).Error
	if err != nil {
		return nil, errx.NewCodeError(errx.NotFundError)
	}

	cInfo := new(types.Menu)
	helper.Swap(info, &cInfo)
	return cInfo, nil
}
