package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Sochi115/todo_cli/app"
)

var appInstance app.App

func main() {
	if len(os.Args) < 2 {
		printTable()
	} else {
		handleArgs()
	}
}

func handleArgs() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTaskName := addCmd.String("task", "", "Adds task")
	addIsHigh := addCmd.Bool("isHigh", false, "Sets priority of the task")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.Int("id", -1, "Deletes task by id")

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Println("Task: ", *addTaskName)
		fmt.Println("Is High Prio: ", *addIsHigh)

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		fmt.Println("ID: ", *deleteId)

	default:
		printTable()
	}
}

func printTable() {
	appInstance.JsonFileName = "resources/dummyJson.json"
	appInstance.GetTodos()
	appInstance.ParseToTable()
}
