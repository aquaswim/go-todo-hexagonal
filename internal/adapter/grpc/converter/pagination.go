package converter

import (
	"hexagonal-todo/internal/adapter/grpc/pb"
	"hexagonal-todo/internal/core/domain"
)

func ConvertPaginationParam(param *pb.PaginationParam) *domain.PaginationParam {
	out := domain.PaginationParam{
		Skip:  0,
		Limit: 10,
	}
	if param != nil {
		out.Skip = int(param.Skip)
		out.Limit = int(param.Limit)
	}
	return &out
}
