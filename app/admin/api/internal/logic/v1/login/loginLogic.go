package login

import (
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/common/static/cKey"
	"blog/database/model"
	"context"
	"github.com/pkg/errors"
	"strconv"
	"time"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	resp = new(types.LoginRes)
	adminInfo := &model.Admin{}
	err = l.svcCtx.DB.WithContext(l.ctx).
		Where("name = ?", req.Username).
		Preload("RoleInfo").
		First(&adminInfo).Error
	if err != nil || adminInfo.ID < 1 {
		return nil, errors.Wrap(errx.NewCodeError(errx.AdminNotFound), "账户不存在！")
	}

	check := helper.PasswordVerify(req.Password, adminInfo.Password)
	if check != true {
		return nil, errx.NewCodeError(errx.AdminNotFound)
	}

	adminId := strconv.FormatInt(adminInfo.ID, 10)
	tokenKey := cKey.AdminTokenKey + adminId
	refreshTokenKey := cKey.AdminRefreshTokenKey + adminId
	//生成token
	roleIdStr := strconv.FormatInt(int64(adminInfo.RoleID), 10)
	token, err := helper.GenToken(l.svcCtx.Config.JwtSecret, adminId, map[string]string{
		"RoleId": roleIdStr,
	}, time.Duration(cKey.AdminTokenTtl))
	if err != nil {
		return nil, err
	}
	//缓存token
	l.svcCtx.Cache.Setex(tokenKey, token, cKey.AdminTokenTtl)
	//生成refreshToken
	str := adminId + token + time.Now().String() + l.svcCtx.Config.JwtSecret
	refreshToken := helper.Hash256(str)
	//缓存token 用来刷新的
	l.svcCtx.Cache.Setex(refreshTokenKey, refreshToken, cKey.AdminRefreshTokenTtl)

	resp.Token = token
	resp.RefreshToken = refreshToken
	return resp, nil
}
