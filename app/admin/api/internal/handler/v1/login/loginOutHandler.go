package login

import (
	"blog/app/admin/api/internal/logic/v1/login"
	"blog/app/admin/api/internal/svc"
	"blog/common/response"
	"net/http"
)

// 退出登录
func LoginOutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := login.NewLoginOutLogic(r.Context(), svcCtx)
		err := l.LoginOut()
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, nil)
		}
	}
}
