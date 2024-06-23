package domain

import (
	"errors"
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
	ErrCodeForbidden
	ErrCodeUnauthorized
)

type AppError struct {
	Code AppErrorCode
	Err  error
}

func (a AppError) Error() string {
	return fmt.Sprintf("App ErrorCode: %d - %s", a.Code, a.Err)
}

func NewAppErrorString(code AppErrorCode, message string) AppError {
	return AppError{Code: code, Err: errors.New(message)}
}

func FromError(code AppErrorCode, err error) AppError {
	return AppError{
		Code: code,
		Err:  err,
	}
}

func (a AppErrorCode) IsErrEqual(err error) bool {
	var appErr AppError
	if !errors.As(err, &appErr) {
		return false
	}
	return appErr.Code == a
}
