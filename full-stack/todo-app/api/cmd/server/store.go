package main

import (
	"example.com/todo-api/internal/list"
	"example.com/todo-api/internal/todo"
)

type ListStore interface {
	Add(list list.List) (list.List, error)
	Get(uid string) (list.List, error)
	GetAll() ([]list.List, error)
	Update(uid string, list list.List) (list.List, error)
	Delete(uid string) error
}

type TodoStore interface {
	Add(todo todo.Todo) (todo.Todo, error)
	Get(uid string) (todo.Todo, error)
	GetAll() ([]todo.Todo, error)
	Update(uid string, todo todo.Todo) (todo.Todo, error)
	Delete(uid string) error
}
