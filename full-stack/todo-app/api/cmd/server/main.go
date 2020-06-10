package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/todo-api/internal/list"
	"example.com/todo-api/internal/todo"
	"github.com/gorilla/mux"
)

func main() {
	listStore, err := list.NewStore()
	if err != nil {
		fmt.Println("error creating new list store:", err.Error())
		os.Exit(1)
	}

	todoStore, err := todo.NewStore()
	if err != nil {
		fmt.Println("error creating new todo store:", err.Error())
		os.Exit(1)
	}

	handler := Handler{
		ListStore: listStore,
		TodoStore: todoStore,
	}

	router := mux.NewRouter()

	router.HandleFunc("/health", handler.Health)
	router.HandleFunc("/lists", handler.AddList).Methods(http.MethodPost)
	router.HandleFunc("/lists", handler.GetAllLists).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}
