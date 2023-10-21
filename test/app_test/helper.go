package test

import "github.com/Sochi115/todo_cli/app"

func generateExpectedTodoList() []app.TodoItem {
	var firstTodo app.TodoItem

	firstTodo.Task = "Throw the trash"
	firstTodo.Id = 1
	firstTodo.Priority = true

	var secondTodo app.TodoItem

	secondTodo.Task = "Walk bella"
	secondTodo.Id = 19
	secondTodo.Priority = false

	var todoList []app.TodoItem

	todoList = append(todoList, firstTodo)
	todoList = append(todoList, secondTodo)

	return todoList
}
