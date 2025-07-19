package conv

import "errors"

type contextKey string

const (
	CtxUserAgent = contextKey("usser-agent")
)

const (
	MessageSuccess = "Success"
)

var (
	ErrInternalServerError  = errors.New("internal server error")
	ErrNotFound             = errors.New("data not found")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrBadParamInput        = errors.New("given param is not valid")
	ErrWrongEmailOrPassword = errors.New("wrong email/password")
)
