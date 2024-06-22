package repositories

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"hexagonal-todo/internal/adapter/storage/pgsql/helpers"
	"hexagonal-todo/internal/core/domain"
	"hexagonal-todo/internal/core/port"
	"time"
)

type todoRepository struct {
	db *pgxpool.Pool
}

func NewTodoRepo(db *pgxpool.Pool) port.TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (t todoRepository) Find(ctx context.Context, paginationParam *domain.PaginationParam) ([]domain.TodoItem, error) {
	sql, args, err := squirrel.
		Select("t.id, t.title, t.description, t.created_at, t.updated_at").
		From("todos t").
		Limit(uint64(paginationParam.Limit)).
		Offset(uint64(paginationParam.Skip)).
		ToSql()
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}

	rows, err := t.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}

	res, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[todoSchema])
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}
	output := make([]domain.TodoItem, len(res))
	for i, row := range res {
		output[i].Id = row.Id
		output[i].Title = row.Title
		output[i].Description = row.Description.String
	}

	return output, nil
}

func (t todoRepository) FindByID(ctx context.Context, id int) (*domain.TodoItem, error) {
	sql, args, err := squirrel.Select("t.id, t.title, t.description, t.created_at, t.updated_at").
		From("todos t").
		Where(squirrel.Eq{"t.id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}

	rows, err := t.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}
	todoRow, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByNameLax[todoSchema])
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}

	return &domain.TodoItem{
		Id:          todoRow.Id,
		Title:       todoRow.Title,
		Description: todoRow.Description.String,
	}, nil
}

func (t todoRepository) Count(ctx context.Context) (int, error) {
	sql, args, err := squirrel.Select("count(t.id)").From("todos t").ToSql()
	if err != nil {
		return 0, helpers.ConvertPgxErrorToAppError(err)
	}

	var count int
	err = t.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		return 0, helpers.ConvertPgxErrorToAppError(err)
	}
	return count, nil
}

func (t todoRepository) Create(ctx context.Context, todo *domain.TodoItem) (*domain.TodoItem, error) {
	sql, args, err := squirrel.Insert("todos").
		Columns("title", "description", "created_at", "updated_at").
		Values(todo.Title, todo.Description, time.Now(), time.Now()).
		Suffix(`RETURNING "id"`).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}

	var id int64
	err = t.db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}
	todo.Id = id
	return todo, nil
}

func (t todoRepository) UpdateByID(ctx context.Context, id int, todo *domain.TodoItem) (*domain.TodoItem, error) {
	sql, args, err := squirrel.Update("todos").
		Set("title", todo.Title).
		Set("description", todo.Description).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}

	_, err = t.db.Exec(ctx, sql, args...)
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}
	todo.Id = int64(id)
	return todo, nil
}

func (t todoRepository) DeleteByID(ctx context.Context, id int) error {
	sql, args, err := squirrel.Delete("todos t").
		Where(squirrel.Eq{"t.id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return helpers.ConvertPgxErrorToAppError(err)
	}

	_, err = t.db.Exec(ctx, sql, args...)
	if err != nil {
		return helpers.ConvertPgxErrorToAppError(err)
	}
	return nil
}
