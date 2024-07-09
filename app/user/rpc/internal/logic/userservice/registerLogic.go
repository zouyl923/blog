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

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *rpc.RegisterReq) (*rpc.RegisterRes, error) {
	password, _ := helper.PasswordHash(in.Password)
	info := model.User{
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Password: password,
	}
	l.svcCtx.DB.Where("phone = ?", in.Phone).First(&info)
	if info.ID > 0 {
		return nil, errors.New("账号已存在！")
	}
	err := l.svcCtx.DB.Create(&info).Error
	if err != nil {
		return nil, err
	}

	userId := strconv.FormatInt(info.ID, 10)
	verify, err := l.svcCtx.VerifyRpc.GenToken(l.ctx, &verifyservice.GenTokenReq{
		Server: "user",
		Key:    userId,
	})
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	var cInfo rpc.User
	helper.ExchangeStruct(info, &cInfo)

	resp := rpc.RegisterRes{}
	resp.Token = verify.Token
	resp.User = &cInfo
	resp.RefreshToken = verify.RefreshToken
	return &resp, nil
}
