package graph


				// This file will be automatically regenerated based on the schema, any resolver implementations
				// will be copied through when generating and any unknown code will be moved to the end.

import (
"context"
"fmt"
"io"
"strconv"
"time"
"sync"
"errors"
"bytes"
gqlparser "github.com/vektah/gqlparser/v2"
"github.com/vektah/gqlparser/v2/ast"
"github.com/99designs/gqlgen/graphql"
"github.com/99designs/gqlgen/graphql/introspection"
"toolbox.com/full-stack-todo/graph/generated"
"toolbox.com/full-stack-todo/graph/model")


















func (r *mutationResolver) CreateUser(ctx context.Context, user model.NewUser) (*model.User, error) {
		newUser, err := r.userStore.Create(user)

	return &newUser, err
	}

func (r *mutationResolver) CreateTodo(ctx context.Context, todo model.NewTodo) (*model.Todo, error) {
		newTodo, err := r.todoStore.Create(todo)

	return &newTodo, err
	}

func (r *mutationResolver) FinishTodo(ctx context.Context, todoID int) (bool, error) {
		err := r.todoStore.Finish(todoID)
	if err != nil {
		return false, err
	}

	return true, nil
	}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID int) (bool, error) {
		err := r.userStore.Delete(userID)
	if err != nil {
		return false, err
	}

	return true, nil
	}

func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID int) (bool, error) {
		err := r.todoStore.Delete(todoID)
	if err != nil {
		return false, err
	}

	return true, nil
	}

func (r *queryResolver) Users(ctx context.Context) ([]model.User, error) {
		panic(fmt.Errorf("not implemented"))
	}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
		panic(fmt.Errorf("not implemented"))
	}

func (r *queryResolver) Todos(ctx context.Context) ([]model.Todo, error) {
		panic(fmt.Errorf("not implemented"))
	}

func (r *queryResolver) Todo(ctx context.Context, id int) (*model.Todo, error) {
		panic(fmt.Errorf("not implemented"))
	}

// Mutation returns generated.MutationResolver implementation.
	func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
// Query returns generated.QueryResolver implementation.
	func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }


type mutationResolver struct { *Resolver }
type queryResolver struct { *Resolver }



    // !!! WARNING !!!
    // The code below was going to be deleted when updating resolvers. It has been copied here so you have
    // one last chance to move it out of harms way if you want. There are two reasons this happens:
	//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
	//    it when you're done.
	//  - You have helper methods in this file. Move them out to keep these resolver files clean.
	var users
var todos
return []model.User{
		{
			ID:   1,
			Name: "Stephen",
		},
		{
			ID:   2,
			Name: "David",
		},
	}, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]model.Todo, error) {
	return []model.Todo{

	}, nil
}

func (r *queryResolver) Todo(ctx context.Context) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

