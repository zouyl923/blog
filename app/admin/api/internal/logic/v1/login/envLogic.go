package login

import (
	"blog/common/helper"
	"context"
	"fmt"
	"net/http"
	"time"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取环境信息
func NewEnvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnvLogic {
	return &EnvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnvLogic) Env(r *http.Request) (resp *types.EnvRes, err error) {
	timestamps := int(time.Now().Unix())
	agent := r.UserAgent()
	ip := r.RemoteAddr

	str := agent + ip + fmt.Sprintf("%x", timestamps)
	uid := helper.Hash256(str)

	resp = new(types.EnvRes)
	resp.Env = "local"
	resp.Uid = uid
	resp.TimeZone = time.Local.String()
	resp.Timestamps = timestamps
	resp.Agent = agent
	resp.Ip = ip
	return resp, nil
}
