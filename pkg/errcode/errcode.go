package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int
	msg     string
	details []string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("error code %d already exists, please change one", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code: %d, error msg: %s", e.Code, e.Msg)
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}
	return e
}

func (e *Error) StatusCode() int {
	switch e.code {
	case Success.code:
		return http.StatusOK
	case ServerError.code:
		return http.StatusInternalServerError
	case InvalidParams.code:
		return http.StatusBadRequest
	case NotFound.code:
		return http.StatusNotFound
	case UnauthorizedTokenError.code, UnauthorizedAuthNotExist.code,
		UnauthorizedTokenGenerate.code:
		return http.StatusUnauthorized
	case UnauthorizedTokenTimeout.code:
		return http.StatusRequestTimeout
	case TooManyRequests.code:
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
