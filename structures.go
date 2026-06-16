package main

import (
	"time"
)

type TodoList struct {
	todos    []Todo
	filename string
}

type Todo struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
