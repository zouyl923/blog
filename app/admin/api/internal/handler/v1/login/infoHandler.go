package login

import (
	"blog/app/admin/api/internal/logic/v1/login"
	"blog/app/admin/api/internal/svc"
	"blog/common/response"
	"net/http"
)

// 管理员信息
func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := login.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info()
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
