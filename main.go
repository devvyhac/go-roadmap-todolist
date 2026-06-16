package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		Usage()
		return
	}

	cmd := os.Args[1]

	tl := &TodoList{filename: "todo.json"}
	file, err := os.OpenFile(tl.filename, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	tl.FetchTodo()

	switch cmd {
	case "list":
		if len(os.Args) == 3 {
			status := []string{"done", "todo", "in-progress"}
			sub := os.Args[2]
			if !slices.Contains(status, sub) {
				Usage()
				return
			}

			tl.ListTodo(sub)
		} else {
			tl.ListTodo("")
		}
	case "add":
		var task string
		if len(os.Args) > 2 {
			task = os.Args[2]
		}
		tl.Add(task)
	case "delete":
		id := ParseID(os.Args[1])
		tl.Delete(id)
	case "update":
		if len(os.Args) != 4 {
			Usage()
			return
		}

		id := ParseID(os.Args[2])
		task := os.Args[3]

		tl.Update(id, task)
	case "mark-in-progress":
		id := ParseID(os.Args[2])
		tl.ChangeStatus(id, "in-progress")
	case "mark-done":
		id := ParseID(os.Args[2])
		tl.ChangeStatus(id, "done")
	default:
		Usage()
		return
	}
}

func ParseID(idStr string) int {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Invalid ID type passed, try again.")
		os.Exit(1)
	}
	return id
}

func Usage() {
	fmt.Println(`Usage:
  task-cli <command> [arguments]

Commands:
  add <description>          Add a new task
  update <id> <description>  Update an existing task's description
  delete <id>                Delete a task
  mark-in-progress <id>      Mark a task status as "in-progress"
  mark-done <id>             Mark a task status as "done"
  list [status]              List tasks (Optional status: todo, in-progress, done)

Examples:
  task-cli add "Buy groceries"
  task-cli update 1 "Buy groceries and cook dinner"
  task-cli delete 1
  task-cli mark-in-progress 1
  task-cli mark-done 1
  task-cli list
  task-cli list done`)
}
