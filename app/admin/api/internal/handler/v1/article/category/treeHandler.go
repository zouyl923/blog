package category

import (
	"blog/app/admin/api/internal/logic/v1/article/category"
	"blog/app/admin/api/internal/svc"
	"blog/common/response"
	"net/http"
)

func TreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewTreeLogic(r.Context(), svcCtx)
		resp, err := l.Tree()
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
