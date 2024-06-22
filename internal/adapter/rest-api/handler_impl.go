package restApi

import (
	"context"
	"hexagonal-todo/internal/core/domain"
	"hexagonal-todo/internal/core/port"
)

type h struct {
	todoService port.TodoService
}

func (h h) HealthCheck(_ context.Context, _ HealthCheckRequestObject) (HealthCheckResponseObject, error) {
	return HealthCheck200JSONResponse{
		Healthy: true,
	}, nil
}

func (h h) TodoItemList(ctx context.Context, request TodoItemListRequestObject) (TodoItemListResponseObject, error) {
	paginationParam := domain.PaginationParam{
		Limit: 10,
		Skip:  0,
	}
	if request.Params.Limit != nil {
		paginationParam.Limit = *request.Params.Limit
	}
	if request.Params.Skip != nil {
		paginationParam.Skip = *request.Params.Skip
	}
	res, err := h.todoService.List(ctx, &paginationParam)
	if err != nil {
		return nil, err
	}

	result := make([]TodoItemWithId, len(res.Items))
	for i := range res.Items {
		todoItemToDtoWithID(&res.Items[i], &result[i])
	}

	return TodoItemList200JSONResponse{
		Meta: &ListMeta{
			Limit: paginationParam.Limit,
			Skip:  paginationParam.Skip,
			Total: res.Count,
		},
		Result: &result,
	}, err
}

func (h h) TodoItemCreate(ctx context.Context, request TodoItemCreateRequestObject) (TodoItemCreateResponseObject, error) {
	createRequest := &domain.TodoItem{
		Title: request.Body.Title,
	}
	if request.Body.Description != nil {
		createRequest.Description = *request.Body.Description
	}
	res, err := h.todoService.Create(ctx, createRequest)
	if err != nil {
		return nil, err
	}
	return &TodoItemCreate200JSONResponse{
		Description: &res.Description,
		Id:          int(res.Id),
		Title:       res.Title,
	}, err
}

func (h h) TodoItemDeleteById(ctx context.Context, request TodoItemDeleteByIdRequestObject) (TodoItemDeleteByIdResponseObject, error) {
	res, err := h.todoService.DeleteByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &TodoItemDeleteById200JSONResponse{
		Description: &res.Description,
		Id:          int(res.Id),
		Title:       res.Title,
	}, err
}

func (h h) TodoItemGetById(ctx context.Context, request TodoItemGetByIdRequestObject) (TodoItemGetByIdResponseObject, error) {
	res, err := h.todoService.FindByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &TodoItemGetById200JSONResponse{
		Description: &res.Description,
		Id:          int(res.Id),
		Title:       res.Title,
	}, err
}

func (h h) TodoItemUpdateById(ctx context.Context, request TodoItemUpdateByIdRequestObject) (TodoItemUpdateByIdResponseObject, error) {
	updateRequest := domain.TodoItem{
		Title: request.Body.Title,
	}
	if request.Body.Description != nil {
		updateRequest.Description = *request.Body.Description
	}

	res, err := h.todoService.UpdateByID(ctx, request.Id, &updateRequest)
	if err != nil {
		return nil, err
	}
	return &TodoItemUpdateById200JSONResponse{
		Description: &res.Description,
		Id:          int(res.Id),
		Title:       res.Title,
	}, err
}
