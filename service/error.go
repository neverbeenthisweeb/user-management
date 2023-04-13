package service

import "net/http"

var (
	ErrMissingUsername        = ServiceErr{Message: "missing username", Code: http.StatusBadRequest}
	ErrMissingPassword        = ServiceErr{Message: "missing password", Code: http.StatusBadRequest}
	ErrPasswordNotFound       = ServiceErr{Message: "password not found", Code: http.StatusNotFound}
	ErrPasswordHashingFailure = ServiceErr{Message: "password hashing failure", Code: http.StatusInternalServerError}
	ErrPasswordWrong          = ServiceErr{Message: "password is wrong", Code: http.StatusUnauthorized}
)

type ServiceErr struct {
	Message string
	Code    int
}

func (se ServiceErr) Error() string {
	return se.Message
}
