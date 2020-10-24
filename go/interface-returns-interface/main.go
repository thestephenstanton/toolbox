package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	t := getData()

	fmt.Println(t)
}

func getData() todo {
	client := http.Client{}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	todo, err := New(resp.Body)
	if err != nil {
		panic(err)
	}

	return todo
}

func New(r io.Reader) (todo, error) {
	var t todo

	err := json.NewDecoder(r).Decode(&t)
	if err != nil {
		return todo{}, nil
	}

	return t, nil
}

type todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
