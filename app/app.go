package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Sochi115/todo_cli/models"
)

type App struct {
	TodoList     []models.TodoItem `json:"todolist"`
	JsonFileName string
}

func (a *App) GetTodos() {
	jsonFile, err := os.Open(a.JsonFileName)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println("NO TASKS")
	} else {
		byteValue, _ := io.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &a)
	}
}

func (a *App) AddTodos(task string) {
	taskItem := constructNewTodoItem(task)

	a.GetTodos()

	a.TodoList = append(a.TodoList, taskItem)
	a.saveTodoListToJson()
}

// func updateJson(todoList models.TodoList) {
// }

func constructNewTodoItem(task string) models.TodoItem {
	var taskItem models.TodoItem

	taskItem.Task = task
	taskItem.SetInitDate()
	taskItem.Priority = true

	return taskItem
}

func (a *App) saveTodoListToJson() {
	file, _ := json.MarshalIndent(a.TodoList, "", "  ")

	err := os.WriteFile(a.JsonFileName, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
