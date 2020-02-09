package kit

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

type Todo struct {
	UID        string `json:"id"`
	IsFinished bool   `json:"isFinished"`
	Text       string `json:"text"`
}

func NewTodo(text string) Todo {
	return Todo{
		UID:        fmt.Sprintf("td_%s", ksuid.New().String()),
		IsFinished: false,
		Text:       text,
	}
}
