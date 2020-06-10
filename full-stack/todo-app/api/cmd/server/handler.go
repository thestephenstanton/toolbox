package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/todo-api/internal/list"
	"github.com/gorilla/mux"
	"github.com/thestephenstanton/hapi"
)

type Handler struct {
	ListStore ListStore
	TodoStore TodoStore
}

func (Handler) Health(w http.ResponseWriter, r *http.Request) {
	hapi.RespondOK(w, "healthy AF")
}

func (handler Handler) AddList(w http.ResponseWriter, r *http.Request) {
	var list list.List
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		fmt.Println(err.Error())
		hapi.RespondBadRequest(w, "could not decode body")
		return
	}

	newList, err := handler.ListStore.Add(list)
	if err != nil {
		fmt.Println(err.Error())
		hapi.RespondInternalError(w, "something went wrong")
		return
	}

	hapi.RespondOK(w, newList)
}

func (handler Handler) GetAllLists(w http.ResponseWriter, r *http.Request) {
	allLists, err := handler.ListStore.GetAll()
	if err != nil {
		fmt.Println(err.Error())
		hapi.RespondInternalError(w, "something went wrong")
		return
	}

	hapi.RespondOK(w, allLists)
}

func (handler Handler) GetList(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	allLists, err := handler.ListStore.Get(id)
	if err != nil {
		fmt.Println(err.Error())
		hapi.RespondInternalError(w, "something went wrong")
		return
	}

	hapi.RespondOK(w, allLists)
}
