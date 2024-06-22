package restApi

import "hexagonal-todo/internal/core/domain"

func todoItemToDtoWithID(in *domain.TodoItem, out *TodoItemWithId) {
	out.Id = int(in.Id)
	out.Title = in.Title
	out.Description = &in.Description
}
