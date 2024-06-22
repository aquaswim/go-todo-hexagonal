package domain

type TodoItem struct {
	Id          int64
	Title       string
	Description string
}

type TodoItemList struct {
	Items []TodoItem
	Count int
}
