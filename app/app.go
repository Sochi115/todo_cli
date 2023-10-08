package app

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type App struct {
	TodoList     []TodoItem
	JsonFileName string
}

func (a *App) GetTodos() {
	jsonFile, err := os.Open(a.JsonFileName)
	defer jsonFile.Close()

	if err == nil {
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &a.TodoList)
	}
}

func (a *App) AddTodos(task string) {
	taskItem := constructNewTodoItem(task)

	a.GetTodos()

	a.TodoList = append(a.TodoList, taskItem)
	a.saveTodoListToJson()
}

// func updateJson(todoList TodoList) {
// }

func constructNewTodoItem(task string) TodoItem {
	var taskItem TodoItem

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
