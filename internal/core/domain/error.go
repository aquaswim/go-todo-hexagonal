package domain

import (
	"fmt"
	"strconv"
)

type AppErrorCode int

func (a AppErrorCode) String() string {
	return strconv.Itoa(int(a))
}

const (
	ErrCodeInternal AppErrorCode = iota + 1
	ErrCodeNotFound
	ErrCodeBadRequest
)

type AppError struct {
	Code AppErrorCode
	Err  error
}

func (a AppError) Error() string {
	return fmt.Sprintf("App ErrorCode: %d - %s", a.Code, a.Err)
}

func FromError(code AppErrorCode, err error) AppError {
	return AppError{
		Code: code,
		Err:  err,
	}
}
