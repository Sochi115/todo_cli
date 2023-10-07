package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Sochi115/todo_cli/models"
)

var TodoList models.TodoList

var jsonFileName string = "resources/dummyJson.json"

func GetTodos(jsonFileName string) {
	jsonFile, err := os.Open(jsonFileName)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println("NO TASKS")
	} else {
		byteValue, _ := io.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &TodoList)
	}
}

func AddTodos(task string) {
	taskItem := constructNewTodoItem(task)

	GetTodos(jsonFileName)

	TodoList.TodoList = append(TodoList.TodoList, taskItem)
	saveTodoListToJson()
}

func updateJson(todoList models.TodoList) {
}

func constructNewTodoItem(task string) models.TodoItem {
	var taskItem models.TodoItem

	taskItem.Task = task
	taskItem.SetInitDate()
	taskItem.Priority = true

	return taskItem
}

func saveTodoListToJson() {
	file, _ := json.MarshalIndent(TodoList, "", "  ")

	err := os.WriteFile(jsonFileName, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
