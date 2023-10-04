package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Sochi115/todo_cli/models"
)

func GetTodos() {
	jsonFileName := "resources/dummyJson.json"

	jsonFile, err := os.Open(jsonFileName)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println("NO TASKS")
	} else {
		byteValue, _ := io.ReadAll(jsonFile)

		var todos models.TodoList

		json.Unmarshal(byteValue, &todos)

		for _, v := range todos.TodoList {
			fmt.Println(v.Task)
			fmt.Println(v.Priority)
			fmt.Println(v.InitDate)
			fmt.Println(v.DueDate)
		}
	}
}
