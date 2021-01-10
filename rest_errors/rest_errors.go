package rest_errors

import (
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}
type restErr struct {
	ErrMessage string `json:"message"`
	ErrStatus int `json:"status"`
	ErrError string `json:"error"`
	ErrCauses []interface{} `json:"causes"`
}

func (r restErr) Message() string {
	return r.ErrMessage
}
func (r restErr) Status() int {
	return r.ErrStatus
}

func (r restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		r.ErrMessage, r.ErrStatus, r.ErrError, r.ErrCauses)
}
func (r restErr) Causes() []interface{} {
	return r.ErrCauses
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

func NewInternalServerError(message string, err error) RestErr {
	rErr := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		rErr.ErrCauses = append(rErr.ErrCauses, err.Error())
	}
	return rErr
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

