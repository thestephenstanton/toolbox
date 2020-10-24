package todos

import (
	"toolbox.com/full-stack-todo/graph/model"
)

var idCounter int

type Store struct {
	db map[int]model.Todo
}

func NewStore() Store {
	return Store{
		db: map[int]model.Todo{},
	}
}

func (s *Store) Add(newTodo model.NewTodo) (model.Todo, error) {
	idCounter++
	todo := model.Todo{
		ID:   idCounter,
		Text: newTodo.Text,
	}

	s.db[idCounter] = todo

	return todo, nil
}

func (s Store) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	for _, todo := range s.db {
		todos = append(todos, todo)
	}

	return todos, nil
}
