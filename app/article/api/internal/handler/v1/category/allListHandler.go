package category

import (
	"blog/app/article/api/internal/logic/v1/category"
	"blog/app/article/api/internal/svc"
	"blog/common/response"
	"net/http"
)

func AllListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewAllListLogic(r.Context(), svcCtx)
		resp, err := l.AllList()
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
