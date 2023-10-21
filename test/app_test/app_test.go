package test

import (
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

	expected := "Test task"

	// Act
	a.AddTodos("Test task", false)
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
}

func TestAddTodoIdReset(t *testing.T) {
	var a *app.App
	a = new(app.App)

	expected := 0

	var initialTask app.TodoItem

	initialTask.Id = 99
	initialTask.Task = "First task"

	a.TodoList = append(a.TodoList, initialTask)

	// Act
	a.AddTodos("Test task", false)
	todoList := a.TodoList

	// Assert
	if len(todoList) <= 1 {
		t.Error("Task not added")
		return
	}
	if len(todoList) >= 3 {
		t.Error("Unexpected tasks in the list")
		t.Log(todoList)
		return
	}

	result := todoList[1].Id
	if result != expected {
		t.Errorf("Got: %v     WANTED: %v", result, expected)
		return
	}
}

func TestDeleteTask(t *testing.T) {
	var a *app.App
	a = new(app.App)
	a.TodoList = generateExpectedTodoList()

	a.RemoveTodoById(19)

	if len(a.TodoList) != 1 {
		t.Error("Failed to delete task 19")
		t.Log(a.TodoList)
	}

	if a.TodoList[0].Task != "Throw the trash" {
		t.Error("Deleted the wrong task")
		t.Log(a.TodoList)
	}
}

func TestDeleteInvalidTaskId(t *testing.T) {
	var a *app.App
	a = new(app.App)
	a.TodoList = generateExpectedTodoList()

	a.RemoveTodoById(100)

	if len(a.TodoList) != 2 {
		t.Error("Task unexpectedly deleted")
		t.Log(a.TodoList)
	}
}
