package main

import (
	"flag"
	"os"

	"github.com/Sochi115/todo_cli/app"
)

var appInstance app.App

func main() {
	appInstance.JsonFileName = "todoJson.json"
	appInstance.GetTodos()
	if len(os.Args) < 2 {
		printTable()
	} else {
		handleArgs()
	}
}

func handleArgs() {
	switch os.Args[1] {
	case "add":
		handleAdd(os.Args[2:])
	case "delete":
		handleDeleteCommand(os.Args[2:])
	case "export":
		handleExportCommand(os.Args[2:])
	default:
		printTable()
	}
	saveChanges()
}

func handleAdd(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTaskName := addCmd.String("task", "", "Adds task")
	addIsHigh := addCmd.Bool("isHigh", false, "Sets priority of the task")

	addCmd.Parse(args)

	appInstance.AddTodos(*addTaskName, *addIsHigh)
	printTable()
}

func handleDeleteCommand(args []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.Int("id", -1, "Deletes task by id")
	deleteCmd.Parse(args)

	appInstance.RemoveTodoById(*deleteId)
	printTable()
}

func handleExportCommand(args []string) {
	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
	filename := exportCmd.String("filename", "", "Relative path to file")
	exportCmd.Parse(args)

	appInstance.ExportAsFile(*filename)
}

func saveChanges() {
	appInstance.SaveTodoListToJson()
}

func printTable() {
	appInstance.ParseToTable()
}
