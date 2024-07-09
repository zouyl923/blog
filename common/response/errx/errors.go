package errx

import "fmt"

type CodeError struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *CodeError) Error() string {

	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.Code, e.Message)
}

func NewError(code int, message string) *CodeError {
	return &CodeError{Code: code, Message: message}
}

func NewCodeError(code int) *CodeError {
	return &CodeError{Code: code, Message: GetCnMessage(code)}
}

func NewMessageError(message string) *CodeError {
	return &CodeError{Code: MessageError, Message: message}
}

func NewSuccess(data interface{}) *CodeError {
	return &CodeError{Code: 0, Message: "ok", Data: data}
}

func NewParamError(message string) *CodeError {
	return &CodeError{Code: ParamError, Message: message}
}
