package port

import (
	"context"
	"hexagonal-todo/internal/core/domain"
)

type TodoRepository interface {
	Find(ctx context.Context, paginationParam *domain.PaginationParam) ([]domain.TodoItem, error)
	FindByID(ctx context.Context, id int) (*domain.TodoItem, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, todo *domain.TodoItem) (*domain.TodoItem, error)
	UpdateByID(ctx context.Context, id int, todo *domain.TodoItem) (*domain.TodoItem, error)
	DeleteByID(ctx context.Context, id int) error
}

type TodoService interface {
	List(ctx context.Context, paginationParam *domain.PaginationParam) (*domain.TodoItemList, error)
	Create(ctx context.Context, todo *domain.TodoItem) (*domain.TodoItem, error)
	FindByID(ctx context.Context, id int) (*domain.TodoItem, error)
	UpdateByID(ctx context.Context, id int, todo *domain.TodoItem) (*domain.TodoItem, error)
	DeleteByID(ctx context.Context, id int) (*domain.TodoItem, error)
}
