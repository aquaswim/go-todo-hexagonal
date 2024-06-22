package service

import (
	"context"
	"hexagonal-todo/internal/core/domain"
	"hexagonal-todo/internal/core/port"
)

type todoService struct {
	todoRepo port.TodoRepository
}

func NewTodoService(todoRepo port.TodoRepository) port.TodoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}

func (t todoService) List(ctx context.Context, paginationParam *domain.PaginationParam) (*domain.TodoItemList, error) {
	result, err := t.todoRepo.Find(ctx, paginationParam)
	if err != nil {
		// todo domain.error
		return nil, err
	}
	count, err := t.todoRepo.Count(ctx)
	if err != nil {
		// todo domain.error
		return nil, err
	}
	return &domain.TodoItemList{
		Items: result,
		Count: count,
	}, nil
}

func (t todoService) Create(ctx context.Context, todo *domain.TodoItem) (*domain.TodoItem, error) {
	createdItem, err := t.todoRepo.Create(ctx, todo)
	if err != nil {
		// todo domain.error
		return nil, err
	}
	return createdItem, nil
}

func (t todoService) FindByID(ctx context.Context, id int) (*domain.TodoItem, error) {
	todoItem, err := t.todoRepo.FindByID(ctx, id)
	if err != nil {
		// todo domain.error
		return nil, err
	}
	return todoItem, nil
}

func (t todoService) UpdateByID(ctx context.Context, id int, todo *domain.TodoItem) (*domain.TodoItem, error) {
	updatedItem, err := t.todoRepo.UpdateByID(ctx, id, todo)
	if err != nil {
		// todo domain.error
		return nil, err
	}
	return updatedItem, err
}

func (t todoService) DeleteByID(ctx context.Context, id int) (*domain.TodoItem, error) {
	todoItem, err := t.FindByID(ctx, id)
	if err != nil {
		// todo domain.error
		return nil, err
	}

	err = t.todoRepo.DeleteByID(ctx, id)
	if err != nil {
		// todo domain.error
		return nil, err
	}
	return todoItem, nil
}
