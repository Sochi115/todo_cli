package app

import (
	"time"
)

type TodoItem struct {
	Id       int    `json:"id"`
	Task     string `json:"task"`
	Priority bool   `json:"priority"`
	InitDate string `json:"init_date"`
	DueDate  string `json:"due_date"`
}

var formatString string = "02-01-2006"

func (t *TodoItem) TogglePriority() {
	t.Priority = !t.Priority
}

func (t *TodoItem) SetInitDate() {
	t.InitDate = time.Now().Format(formatString)
}

func (t *TodoItem) SetDueDate(days int) {
	curr := time.Now().AddDate(0, 0, days)
	t.DueDate = curr.Format(formatString)
}
