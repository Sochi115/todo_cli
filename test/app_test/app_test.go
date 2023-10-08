package test

import (
	"os"
	"testing"

	"github.com/Sochi115/todo_cli/app"
)

func TestGetTodosWithFileName(t *testing.T) {
	// Arrange
	expected := generateExpectedTodoList()
	var a *app.App
	a = new(app.App)
	a.JsonFileName = "../test_resources/dummyJson.json"

	// Act
	a.GetTodos()
	result := a.TodoList

	// Assert
	if len(result) == 0 {
		t.Error("File not found")
		return
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Got: %v     WANTED: %v", result[i], expected[i])
			return
		}
	}
}

func TestAddTodo(t *testing.T) {
	// Arrange
	var a *app.App
	a = new(app.App)
	a.JsonFileName = "../test_resources/dummyAddJson.json"

	expected := "Test task"

	// Act
	a.AddTodos("Test task")
	todoList := a.TodoList

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
	os.Remove(a.JsonFileName)
}

func generateExpectedTodoList() []app.TodoItem {
	var firstTodo app.TodoItem

	firstTodo.Task = "Throw the trash"
	firstTodo.Priority = true
	firstTodo.InitDate = "11-05-2000"

	var secondTodo app.TodoItem

	secondTodo.Task = "Walk bella"
	secondTodo.Priority = false
	secondTodo.InitDate = "11-05-2000"
	secondTodo.DueDate = "12-05-2000"

	var todoList []app.TodoItem

	todoList = append(todoList, firstTodo)
	todoList = append(todoList, secondTodo)

	return todoList
}
