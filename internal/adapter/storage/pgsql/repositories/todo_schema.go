package repositories

import (
	"database/sql"
)

type todoSchema struct {
	Id          int64          `db:"id"`
	Title       string         `db:"title"`
	Description sql.NullString `db:"description"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
}
