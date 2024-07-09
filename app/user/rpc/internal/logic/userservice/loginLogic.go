package userservicelogic

import (
	"blog/app/verify/rpc/client/verifyservice"
	"blog/common/helper"
	"blog/database/model"
	"context"
	"github.com/pkg/errors"
	"strconv"

	"blog/app/user/rpc/internal/svc"
	"blog/app/user/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *rpc.LoginReq) (*rpc.LoginRes, error) {
	info := model.User{}
	err := l.svcCtx.DB.Where("phone = ?", in.Phone).First(&info).Error
	if err != nil {
		return nil, err
	}
	if info.ID < 1 {
		return nil, errors.New("账号或者密码错误！")
	}
	check := helper.PasswordVerify(in.Password, info.Password)
	if !check {
		return nil, errors.New("账号或者密码错误！")
	}

	userId := strconv.FormatInt(info.ID, 10)
	verify, err := l.svcCtx.VerifyRpc.GenToken(l.ctx, &verifyservice.GenTokenReq{
		Server: "user",
		Key:    userId,
	})
	var cInfo rpc.User
	helper.ExchangeStruct(info, &cInfo)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	return &rpc.LoginRes{
		Token:        verify.Token,
		RefreshToken: verify.RefreshToken,
		User:         &cInfo,
	}, nil
}
