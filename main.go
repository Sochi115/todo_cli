package main

import (
	"fmt"
	"time"

	"github.com/Sochi115/todo_cli/app"
)

func main() {
	start := time.Now()
	// app.GetTodos()
	var app app.App
	// app.JsonFileName = "resources/dummyJson.json"
	app.JsonFileName = "dummyJson.json"

	app.GetTodos()
	app.ParseToTable()
	fmt.Println(time.Since(start))
}
