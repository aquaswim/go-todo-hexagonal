package helpers

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"hexagonal-todo/internal/core/domain"
)

func ConvertPgxErrorToAppError(err error) domain.AppError {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return domain.FromError(domain.ErrCodeNotFound, err)
	default:
		return domain.FromError(domain.ErrCodeInternal, err)
	}
}
