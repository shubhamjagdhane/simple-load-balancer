package errors

import (
	"fmt"
)

// defines the error keys and messages
const (
	BadRequest         = "Bad Request"
	ServiceUnavailable = "Service Unavailable"
	BadContentType     = "Bad Content Type"
	BadHeader          = "Bad Header"
)

type IError interface {
	Error() string
}

func (o CustomError) Error() string {
	return fmt.Sprintf(`{"http_code":%v, "error":"%v"}`, o.HTTPCode, o.Message)
}

type CustomError struct {
	Message  string `json:"message"`
	HTTPCode int    `json:"status_code"`
}

func GetError(errMsg string, statusCode int) *CustomError {
	return &CustomError{
		Message:  errMsg,
		HTTPCode: statusCode,
	}
}
