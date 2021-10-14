package utils

import (
	"fmt"
)

var ApiErr ApiError

func CreateApiErr(code int64, msg string) *ApiError {
	return &ApiError{Code: code, Msg: msg}
}

type ApiError struct {
	Code int64
	Msg  string
}

func (ae *ApiError) Error() string {

	return fmt.Sprintf("%v@%s", ae.Code, ae.Msg)

}
