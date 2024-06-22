package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"hexagonal-todo/internal/adapter/storage/pgsql/helpers"
	"hexagonal-todo/internal/core/domain"
	"hexagonal-todo/internal/core/port"
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
	rows, err := t.db.Query(ctx, `select t.id, t.title, t.description, t.created_at, t.updated_at from todos t order by t.id limit $1 offset $2`, paginationParam.Limit, paginationParam.Skip)
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
	rows, err := t.db.Query(ctx, "select t.id, t.title, t.description, t.created_at, t.updated_at from todos t where t.id = $1", id)
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
	var count int
	err := t.db.QueryRow(ctx, "select count(t.id) from todos t").Scan(&count)
	if err != nil {
		return 0, helpers.ConvertPgxErrorToAppError(err)
	}
	return count, nil
}

func (t todoRepository) Create(ctx context.Context, todo *domain.TodoItem) (*domain.TodoItem, error) {
	var id int64
	err := t.db.QueryRow(ctx, `INSERT INTO todos
(title, description, created_at, updated_at)
VALUES($1, $2, now(), now()) RETURNING id`, todo.Title, todo.Description).Scan(&id)
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}
	todo.Id = id
	return todo, nil
}

func (t todoRepository) UpdateByID(ctx context.Context, id int, todo *domain.TodoItem) (*domain.TodoItem, error) {
	_, err := t.db.Exec(ctx, `UPDATE todos
	SET title=$2,description=$3,updated_at='now()'
	WHERE id=$1
`, id, todo.Title, todo.Description)
	if err != nil {
		return nil, helpers.ConvertPgxErrorToAppError(err)
	}
	todo.Id = int64(id)
	return todo, nil
}

func (t todoRepository) DeleteByID(ctx context.Context, id int) error {
	_, err := t.db.Exec(ctx, `DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return helpers.ConvertPgxErrorToAppError(err)
	}
	return nil
}
