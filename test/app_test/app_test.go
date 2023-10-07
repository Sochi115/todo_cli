package test

import (
	"testing"

	"github.com/Sochi115/todo_cli/app"
	"github.com/Sochi115/todo_cli/models"
)

func TestGetTodosWithFileName(t *testing.T) {
	// Arrange
	resetTodoList()
	expected := generateExpectedTodoList()
	jsonFileName := "../test_resources/dummyJson.json"

	// Act
	app.GetTodos(jsonFileName)
	result := app.TodoList.TodoList

	// Assert
	if len(result) == 0 {
		t.Error("File not found")
		return
	}
	for i := range result {
		if result[i] != expected.TodoList[i] {
			t.Errorf("Got: %v     WANTED: %v", result[i], expected.TodoList[i])
			return
		}
	}
}

func TestAddTodo(t *testing.T) {
	// Arrange
	resetTodoList()
	expected := "Test task"

	// Act
	app.AddTodos("Test task")
	todoList := app.TodoList.TodoList

	// Assert
	if len(todoList) == 0 {
		t.Error("Task not added")
		return
	}
	if len(todoList) >= 2 {
		t.Error("Unexpected tasks in the list")
		t.Log(todoList)
		return
	}

	result := todoList[0].Task
	if result != expected {
		t.Errorf("Got: %v     WANTED: %v", result, expected)
		return
	}
}

func generateExpectedTodoList() models.TodoList {
	var firstTodo models.TodoItem

	firstTodo.Task = "Throw the trash"
	firstTodo.Priority = true
	firstTodo.InitDate = "11-05-2000"

	var secondTodo models.TodoItem

	secondTodo.Task = "Walk bella"
	secondTodo.Priority = false
	secondTodo.InitDate = "11-05-2000"
	secondTodo.DueDate = "12-05-2000"

	var todoList models.TodoList

	todoList.TodoList = append(todoList.TodoList, firstTodo)
	todoList.TodoList = append(todoList.TodoList, secondTodo)

	return todoList
}

func resetTodoList() {
	app.TodoList.TodoList = []models.TodoItem{}
}
