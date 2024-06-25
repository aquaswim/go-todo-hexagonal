package converter

import "hexagonal-todo/internal/adapter/grpc/pb"

func NewListMeta(limit, skip, total int) *pb.ListMeta {
	return &pb.ListMeta{
		Limit: int64(limit),
		Skip:  int64(skip),
		Total: int64(total),
	}
}
