package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"hexagonal-todo/internal/adapter/grpc/converter"
	"hexagonal-todo/internal/adapter/grpc/pb"
	"hexagonal-todo/internal/core/domain"
	"hexagonal-todo/internal/core/port"
)

type handlerWithAuth struct {
	pb.UnimplementedTodoHexagonalServiceWithAuthServer
	authService port.AuthService `container:"type"`
	todoService port.TodoService `container:"type"`
}

func (h handlerWithAuth) AuthProfile(ctx context.Context, _ *emptypb.Empty) (*pb.Auth_UserData, error) {
	profile, err := h.authService.MyProfile(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.Auth_UserData{
		Id:       profile.Id,
		Email:    profile.Email,
		Password: "-redacted-",
		FullName: profile.FullName,
	}, nil
}

func (h handlerWithAuth) TodoFind(ctx context.Context, param *pb.Todo_FindPayload) (*pb.Todo_ListResult, error) {
	paginationParam := converter.ConvertPaginationParam(param.GetPagination())
	res, err := h.todoService.List(ctx, paginationParam)
	if err != nil {
		return nil, err
	}

	items := make([]*pb.Todo_Data, len(res.Items))
	for i, todoItem := range res.Items {
		items[i] = &pb.Todo_Data{
			Id:          todoItem.Id,
			Title:       todoItem.Title,
			Description: todoItem.Description,
		}
	}

	return &pb.Todo_ListResult{
		Items: items,
		Meta:  converter.NewListMeta(paginationParam.Limit, paginationParam.Skip, res.Count),
	}, nil
}

func (h handlerWithAuth) TodoGetByID(ctx context.Context, payload *pb.Todo_GetByIDPayload) (*pb.Todo_Data, error) {
	res, err := h.todoService.FindByID(ctx, int(payload.Id))

	if err != nil {
		return nil, err
	}

	return &pb.Todo_Data{
		Id:          res.Id,
		Title:       res.Title,
		Description: res.Description,
	}, nil
}

func (h handlerWithAuth) TodoCreate(ctx context.Context, payload *pb.Todo_CreatePayload) (*pb.Todo_Data, error) {
	res, err := h.todoService.Create(ctx, &domain.TodoItem{
		Title:       payload.Title,
		Description: payload.Description,
	})

	if err != nil {
		return nil, err
	}

	return &pb.Todo_Data{
		Id:          res.Id,
		Title:       res.Title,
		Description: res.Description,
	}, nil
}

func (h handlerWithAuth) TodoUpdateByID(ctx context.Context, payload *pb.Todo_UpdatePayload) (*pb.Todo_Data, error) {
	res, err := h.todoService.UpdateByID(ctx, int(payload.FindPayload.Id), &domain.TodoItem{
		Title:       payload.UpdatePayload.Title,
		Description: payload.UpdatePayload.Description,
	})

	if err != nil {
		return nil, err
	}

	return &pb.Todo_Data{
		Id:          res.Id,
		Title:       res.Title,
		Description: res.Description,
	}, nil
}

func (h handlerWithAuth) TodoDeleteByID(ctx context.Context, payload *pb.Todo_GetByIDPayload) (*pb.Todo_Data, error) {
	res, err := h.todoService.DeleteByID(ctx, int(payload.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Todo_Data{
		Id:          res.Id,
		Title:       res.Title,
		Description: res.Description,
	}, nil
}
