package login

import (
	"blog/app/admin/api/internal/logic/v1/login"
	"blog/app/admin/api/internal/svc"
	"blog/common/response"
	"net/http"
)

// 获取环境信息
func EnvHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := login.NewEnvLogic(r.Context(), svcCtx)
		resp, err := l.Env(r)
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
