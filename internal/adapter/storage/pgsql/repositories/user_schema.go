package repositories

import (
	"database/sql"
	"hexagonal-todo/internal/core/domain"
)

type userSchema struct {
	Id        int64          `db:"id"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	FullName  sql.NullString `db:"full_name"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}

func (u userSchema) toUserData() *domain.UserDataWithID {
	return &domain.UserDataWithID{
		Id: u.Id,
		UserData: domain.UserData{
			LoginCredential: domain.LoginCredential{
				Email:    u.Email,
				Password: u.Password,
			},
			FullName: u.FullName.String,
		},
	}
}
