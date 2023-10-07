package main

import (
	"fmt"
	"time"

	"github.com/Sochi115/todo_cli/app"
)

func main() {
	start := time.Now()
	// app.GetTodos()
	var jsonFileName string = "resources/dummyJson.json"
	app.GetTodos(jsonFileName)
	app.ParseToTable(app.TodoList)
	fmt.Println(time.Since(start))
}
