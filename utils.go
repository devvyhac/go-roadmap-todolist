package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

func (t *TodoList) FetchTodo() []Todo {
	file, err := os.ReadFile(t.filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	json.Unmarshal(file, &t.todos)
	return t.todos
}

func (t *TodoList) ListTodo(status string) {
	if len(t.todos) == 0 {
		fmt.Println("No todos added yet.")
		return
	}

	if status == "" {
		for _, todo := range t.todos {
			fmt.Printf("ID: %s - Status: %s - '%s'\n", todo.Id, todo.Status, todo.Description)
		}
		return
	}
	for _, todo := range t.todos {
		if todo.Status == status {
			fmt.Printf("ID: %s - Status: %s - '%s'\n", todo.Id, todo.Status, todo.Description)
		}
	}
}

func (t *TodoList) Add(task string) {
	if task == "" {
		Usage()
		return
	}
	id := len(t.todos) + 1
	if len(t.todos) > 0 {
		id = t.todos[len(t.todos)-1].Id + 1
	}
	now := time.Now()
	todo := Todo{Id: id, Description: task, Status: "todo", CreatedAt: now, UpdatedAt: now}
	t.todos = append(t.todos, todo)

	t.Save()

	fmt.Printf("Task added successfully (ID: %d)\n", id)
}

func (t *TodoList) Save() {
	data, err := json.MarshalIndent(t.todos, "", "\t")
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := os.WriteFile(t.filename, data, 0644); err != nil {
		log.Fatal(err)
		return
	}
}

func (t *TodoList) Delete(id int) {
	index := TodoIndex(t, id)
	t.todos = append(t.todos[:index], t.todos[index+1:]...)
	t.Save()
	fmt.Printf("Todo task with the id(%d), deleted!\n", id)
}

func (t *TodoList) Update(id int, task string) {
	index := TodoIndex(t, id)

	t.todos[index].Description = task
	t.todos[index].UpdatedAt = time.Now()

	t.Save()
	fmt.Printf("Todo task with the id(%d), updated!\n", id)
}

func (t *TodoList) ChangeStatus(id int, status string) {
	index := TodoIndex(t, id)

	t.todos[index].Status = status
	t.todos[index].UpdatedAt = time.Now()

	t.Save()
	fmt.Printf("Todo ID(%d) is now (%s)\n", id, status)
}

func TodoIndex(t *TodoList, id int) int {
	index := slices.IndexFunc(t.todos, func(todo Todo) bool {
		return todo.Id == id
	})

	if index == -1 {
		fmt.Printf("Todo with the id(%d) doesn't exist!\n", id)
		os.Exit(1)
	}

	return index
}
