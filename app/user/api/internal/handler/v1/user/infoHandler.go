package user

import (
	"blog/app/user/api/internal/logic/v1/user"
	"blog/app/user/api/internal/svc"
	"blog/common/response"
	"net/http"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info()
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
