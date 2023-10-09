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
	a.GetTodos()
	newId := a.generateId()

	taskItem := constructNewTodoItem(task, newId)

	a.TodoList = append(a.TodoList, taskItem)
	a.saveTodoListToJson()
}

func (a *App) generateId() int {
	todoLength := len(a.TodoList)
	if todoLength == 0 {
		return 0
	}

	lastId := a.TodoList[todoLength-1].Id

	if lastId >= 99 {
		return 0
	}

	return lastId + 1
}

func constructNewTodoItem(task string, id int) TodoItem {
	var taskItem TodoItem

	taskItem.Task = task
	taskItem.Id = id
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
