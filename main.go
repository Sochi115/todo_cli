package main

import (
	"fmt"
	"time"

	"github.com/Sochi115/todo_cli/app"
)

func main() {
	start := time.Now()
	app.GetTodos()
	fmt.Println(time.Since(start))
}
