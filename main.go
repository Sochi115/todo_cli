package main

import (
	"fmt"
	"time"
)

type TodoItem struct {
	Task     string
	Priority string
	InitDate time.Time
	DueDate  time.Time
}

func main() {
	var todo TodoItem

	todo.InitDate = time.Now()
	todo.DueDate = time.Now()
	todo.Task = "Test Task"
	todo.Priority = "med"
	fmt.Println(todo)
}
