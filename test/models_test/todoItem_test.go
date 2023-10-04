package test

import (
	"testing"
	"time"

	"github.com/Sochi115/todo_cli/models"
)

func TestSetInitDate(t *testing.T) {
	expected := time.Now().Format("02-01-2006")

	var taskItem models.TodoItem

	taskItem.SetInitDate()

	if expected != taskItem.InitDate {
		t.Errorf("Got: %q     WANTED: %q", taskItem.InitDate, expected)
	}
}

func TestSetDueDate(t *testing.T) {
	formatString := "02-01-2006"
	currDate := time.Now().AddDate(0, 0, 2)
	expected := currDate.Format(formatString)

	var taskItem models.TodoItem

	taskItem.SetDueDate(2)

	if expected != taskItem.DueDate {
		t.Errorf("Got: %q     WANTED: %q", taskItem.DueDate, expected)
	}
}
