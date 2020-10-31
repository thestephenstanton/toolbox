package main

import (
	"time"

	"toolbox.com/dataloader/todos"
)

func main() {
	repo := todos.NewRepo()

	ids := []int{1, 2}
	for _, id := range ids {
		go func(id int) {
			// repo.OldGet(id)
			repo.NewGet(id)
		}(id)
	}

	// time.Sleep(20 * time.Millisecond)

	// ids = []int{1, 2, 3}
	// for _, id := range ids {
	// 	go func(id int) {
	// 		// repo.OldGet(id)
	// 		repo.NewGet(id)
	// 	}(id)
	// }

	time.Sleep(1 * time.Second)
}
