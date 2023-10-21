package app

import (
	"encoding/json"
	"fmt"
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

func (a *App) AddTodos(task string, prio bool) {
	taskItem := a.constructNewTodoItem(task, prio)

	a.TodoList = append(a.TodoList, taskItem)
}

func (a *App) RemoveTodoById(id int) {
	for idx, t := range a.TodoList {
		if t.Id == id {
			a.TodoList = a.removeFromSlice(idx)
			return
		}
	}
}

func (a *App) SaveTodoListToJson() {
	file, _ := json.MarshalIndent(a.TodoList, "", "  ")

	err := os.WriteFile(a.JsonFileName, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) ExportAsFile(filename string) {
	file, _ := json.MarshalIndent(a.TodoList, "", "  ")

	if filename == "" {
		filename = "todo.json"
	}
	filename = ".\\" + filename

	fmt.Println("Writing to file:", filename)
	err := os.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) removeFromSlice(idx int) []TodoItem {
	todoLength := len(a.TodoList) - 1

	if idx < todoLength {
		a.TodoList[idx] = a.TodoList[todoLength]
	}
	return a.TodoList[:todoLength]
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

func (a *App) constructNewTodoItem(task string, prio bool) TodoItem {
	var taskItem TodoItem

	taskItem.Task = task
	taskItem.Id = a.generateId()
	taskItem.Priority = prio

	return taskItem
}
