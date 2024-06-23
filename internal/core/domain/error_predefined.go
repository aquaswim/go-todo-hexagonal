package domain

// this file will save all of this application errors

var (
	AppErrEmailAlreadyExists   = NewAppErrorString(ErrCodeBadRequest, "email already exists")
	AppErrUserNotFound         = NewAppErrorString(ErrCodeNotFound, "user not found")
	AppErrInvalidToken         = NewAppErrorString(ErrCodeForbidden, "invalid token")
	AppErrTokenNotProvided     = NewAppErrorString(ErrCodeUnauthorized, "token not provided")
	AppErrTokenInvalidProvided = NewAppErrorString(ErrCodeUnauthorized, "provided token is invalid")
	AppErrAuthCtxEmpty         = NewAppErrorString(ErrCodeUnauthorized, "user auth context is empty")
)
