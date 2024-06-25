package repositories

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"hexagonal-todo/internal/adapter/storage/pgsql/helpers"
	"hexagonal-todo/internal/core/domain"
	"hexagonal-todo/internal/core/port"
	"time"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) port.UserRepository {
	log.Debug().Msg("initializing user repositories")

	return &userRepository{
		db: db,
	}
}

func (u userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.UserDataWithID, error) {
	sql, args, err := squirrel.Select("u.id", "u.email", "u.password", "u.full_name", "u.created_at", "u.updated_at").
		From("users u").
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}

	rows, err := u.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}
	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByNameLax[userSchema])
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}

	return user.toUserData(), nil
}

func (u userRepository) GetUserById(ctx context.Context, id int64) (*domain.UserDataWithID, error) {
	sql, args, err := squirrel.Select("u.id", "u.email", "u.password", "u.full_name", "u.created_at", "u.updated_at").
		From("users u").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}

	rows, err := u.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}
	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByNameLax[userSchema])
	if err != nil {
		return nil, err
	}

	return user.toUserData(), nil
}

func (u userRepository) CreateUser(ctx context.Context, user *domain.UserData) (*domain.UserDataWithID, error) {
	sql, args, err := squirrel.Insert("users").
		Columns("email", "password", "full_name", "created_at", "updated_at").
		Values(user.Email, user.Password, user.FullName, time.Now(), time.Now()).
		Suffix(`RETURNING "id"`).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}
	var id int64
	err = u.db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}
	return &domain.UserDataWithID{
		Id:       id,
		UserData: *user,
	}, nil
}
