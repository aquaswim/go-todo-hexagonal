package repositories

import (
	"context"
	"fmt"
	"hexagonal-todo/internal/adapter/config"
	"hexagonal-todo/internal/adapter/storage/pgsql"
	"hexagonal-todo/internal/core/domain"
	"testing"
	"time"
)

func newRepo(tb testing.TB) *todoRepository {
	//todo: use https://github.com/pashagolub/pgxmock
	tb.Log("connecting to database")
	dbPool, err := pgsql.Connect(config.DBConfigFromENV())
	if err != nil {
		tb.Fatalf("fail to connect to db: %s", err)
	}
	return &todoRepository{
		db: dbPool,
	}
}

func TestTodoRepository_Find(t *testing.T) {
	ctx := context.TODO()
	repo := newRepo(t)
	_, err := repo.Find(ctx, &domain.PaginationParam{
		Skip:  0,
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("Find return failed: %s", err)
	}
}

func TestTodoRepository_FindByID(t *testing.T) {
	ctx := context.TODO()
	repo := newRepo(t)
	_, err := repo.FindByID(ctx, 1)
	if err != nil {
		t.Fatalf("FindByID return failed: %s", err)
	}
}

func TestTodoRepository_Count(t *testing.T) {
	ctx := context.TODO()
	repo := newRepo(t)
	_, err := repo.Count(ctx)
	if err != nil {
		t.Fatalf("Count return failed: %s", err)
	}
}

func TestTodoRepository_CreateAndDelete(t *testing.T) {
	ctx := context.TODO()
	repo := newRepo(t)
	var createdId int64
	t.Run("Create", func(t *testing.T) {
		createdTodo, err := repo.Create(ctx, &domain.TodoItem{
			Title:       fmt.Sprintf("Title: %s", time.Now()),
			Description: fmt.Sprintf("description: %s", time.Now()),
		})
		if err != nil {
			t.Fatalf("Create return failed: %s", err)
		}
		createdId = createdTodo.Id
	})
	t.Run("Delete", func(t *testing.T) {
		if createdId == 0 {
			t.Skip("create test is failing")
		}
		err := repo.DeleteByID(ctx, int(createdId))
		if err != nil {
			t.Fatalf("Delete return failed: %s", err)
		}
	})
}

func TestTodoRepository_UpdateByID(t *testing.T) {
	ctx := context.TODO()
	repo := newRepo(t)
	_, err := repo.UpdateByID(ctx, 1, &domain.TodoItem{
		Title:       fmt.Sprintf("Updated: %s", time.Now()),
		Description: fmt.Sprintf("Updated: %s", time.Now()),
	})
	if err != nil {
		t.Fatalf("Create return failed: %s", err)
	}
}
