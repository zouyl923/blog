package {{.PkgName}}

import (
	"net/http"
    "blog/common/response"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/locales/zh"
    zhTrans "github.com/go-playground/validator/v10/translations/zh"
    ut "github.com/go-playground/universal-translator"
    "reflect"
	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		//解析参数
        httpx.Parse(r, &req)
        //验证器
        uni := ut.New(zh.New())
        trans, _ := uni.GetTranslator("zh")
        validate := validator.New()
        zhTrans.RegisterDefaultTranslations(validate, trans)
        validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
            return fld.Tag.Get("label")
        })
        errs := validate.Struct(req)
        if errs != nil {
            first := errs.(validator.ValidationErrors)[0]
            response.ParamError(w, first.Translate(trans))
            return
        }

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			response.Error(w, err)
		} else {
			{{if .HasResp}}response.Success(w, resp)
            {{else}}response.Success(w, nil){{end}}
		}
	}
}
