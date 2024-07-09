package response

import (
	"blog/common/response/errx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New() *Response {
	return &Response{}
}

func Success(w http.ResponseWriter, data interface{}) {
	r := &Response{}
	r.Code = errx.Success
	r.Message = errx.GetCnMessage(r.Code)
	r.Data = data
	httpx.OkJson(w, r)
}

func Error(w http.ResponseWriter, err error) {
	causeErr := errors.Cause(err)
	e, ok := causeErr.(*errx.CodeError)

	r := &Response{}
	if ok {
		r.Code = e.Code
		r.Message = e.Message
	} else {
		code := errx.UnKnowError
		r.Code = code
		r.Message = errx.GetCnMessage(code)
	}
	httpx.OkJson(w, r)
}

func ParamError(w http.ResponseWriter, message string) {
	r := &Response{}
	r.Code = errx.ParamError
	r.Message = message
	httpx.OkJson(w, r)
}
