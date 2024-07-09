package role

import (
	"blog/app/admin/api/internal/logic/v1/system/role"
	"blog/app/admin/api/internal/svc"
	"blog/common/response"
	"net/http"
)

func AllListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := role.NewAllListLogic(r.Context(), svcCtx)
		resp, err := l.AllList()
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
