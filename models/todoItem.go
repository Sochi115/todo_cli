package models

import (
	"time"
)

type TodoItem struct {
	Task     string
	Priority bool
	InitDate string
	DueDate  string
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
