package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thestephenstanton/hapi"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	"github.com/thestephenstanton/hapi/errors"
)

type todo struct {
	ID       string `json:"id"`
	Task     string `json:"task"`
	Finished bool   `json:"finished"`
}

var db []*todo

func main() {
	db = []*todo{}

	router := mux.NewRouter()

	router.HandleFunc("/health", healthCheck)
	router.HandleFunc("/todos", getTodos).Methods(http.MethodGet)
	router.HandleFunc("/todos", addTodo).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}", updateTodo).Methods(http.MethodPut)
	router.HandleFunc("/todos/{id}", deleteTodo).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", router)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	err := hapi.RespondOK(w, "healthy AF")
	if err != nil {
		hapi.RespondInternalError(w, err)
		return
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	err := hapi.RespondOK(w, db)
	if err != nil {
		hapi.RespondInternalError(w, err)
		return
	}
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo todo
	err := hapi.UnmarshalBody(r, &newTodo)
	if err != nil {
		hapi.RespondInternalError(w, err)
		return
	}

	json.Unmarshal([]byte{}, newTodo)
	

	fmt.Println(newTodo)

	newTodo.ID = ksuid.New().String()

	db = append(db, &newTodo)

	err = hapi.RespondOK(w, newTodo)
	if err != nil {
		hapi.RespondInternalError(w, err)
		return
	}
}

type updateTodoRequest struct {
	Task     *string `json:"task"`
	Finished *bool   `json:"finished"`
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := hapi.GetQueryParam(r, "id")
	if err != nil {
		hapi.RespondBadRequest(w, "request not in proper format")
		return
	}

	var todoToUpdate *todo
	for _, todo := range db {
		if todo.ID == id {
			todoToUpdate = todo
		}
	}

	if todoToUpdate == nil {
		hapi.RespondBadRequest(w, fmt.Sprintf("could not find index '%s'", id))
		return
	}

	var req updateTodoRequest
	hapi.UnmarshalBody(r, req)
	if err != nil {
		hapi.RespondInternalError(w, err)
		return
	}

	if req.Task != nil {
		todoToUpdate.Task = *req.Task
	}

	if req.Finished != nil {
		todoToUpdate.Finished = *req.Finished
	}

	hapi.RespondOK(w, nil)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var indexToDelete *int
	for i, todo := range db {
		if todo.ID == id {
			indexToDelete = &i
		}
	}

	if indexToDelete == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(errors.Newf("could not find index '%s'", id))
		return
	}

	db = append(db[:*indexToDelete], db[*indexToDelete+1:]...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
