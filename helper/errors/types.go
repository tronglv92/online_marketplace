package errors

import (
	"net/http"
)

const (
	StackKey   = "stack"
	CodeKey    = "code"
	DefaultMsg = "Something went wrong"
)

type Error interface {
	Error() string
	GetCode() int
	GetCause() error
	GetReason() string

	// HasReport() bool
	// GetMetaCode() string
	// GetMeta(key string) string
	// GetMetaData() map[string]string
}

func BadRequest(err error, opts ...Option) Error {
	return New(http.StatusBadRequest, err, opts...)
}
func Unauthorized(err error, opts ...Option) Error {
	return New(http.StatusUnauthorized, err, opts...)
}
func InternalServerReason(reason string, opts ...Option) Error {
	return Newf(http.StatusInternalServerError, reason, opts...)
}

func New(code int, err error, opts ...Option) Error {

	r := &errorSvc{
		Status: Status{
			Code:     code,
			Reason:   err.Error(),
			Metadata: make(map[string]string),
		},
		cause: err,
	}
	for _, applyOpt := range opts {
		if applyOpt == nil {
			continue
		}
		applyOpt(r)
	}
	return r
}

func Newf(code int, reason string, opts ...Option) Error {
	r := &errorSvc{
		Status: Status{
			Code:     code,
			Reason:   reason,
			Metadata: make(map[string]string),
		},
	}
	for _, applyOpt := range opts {
		if applyOpt == nil {
			continue
		}
		applyOpt(r)
	}
	return r
}

func DataNotFound() Error {
	return Newf(http.StatusOK, "No Data")
}

func InternalServer(err error, opts ...Option) Error {
	return New(http.StatusInternalServerError, err, opts...)
}
