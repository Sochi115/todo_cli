package app

type TodoItem struct {
	Id       int    `json:"id"`
	Task     string `json:"task"`
	Priority bool   `json:"priority"`
}

var formatString string = "02-01-2006"

func (t *TodoItem) TogglePriority() {
	t.Priority = !t.Priority
}
