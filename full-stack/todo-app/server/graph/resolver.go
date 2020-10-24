package graph

import "toolbox.com/full-stack-todo/graph/model"

type userStore interface {
	Create(model.NewUser) (model.User, error)
	GetAll() ([]model.User, error)
	Delete(userID int) error
}

type todoStore interface {
	Create(model.NewTodo) (model.Todo, error)
	GetAll() ([]model.Todo, error)
	Finish(todoID int) error
	Delete(todoID int) error
}

func NewResolver(userStore userStore, todoStore todoStore) Resolver {
	return Resolver{
		userStore: userStore,
		todoStore: todoStore,
	}
}

type Resolver struct {
	userStore userStore
	todoStore todoStore
}
