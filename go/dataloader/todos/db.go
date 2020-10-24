package todos

import "fmt"

type db struct {
	data map[int]Todo
}

func NewDB() db {
	return db{
		data: map[int]Todo{
			1: {
				ID:     1,
				Text:   "take out trash",
				IsDone: false,
			},
			2: {
				ID:     2,
				Text:   "walk dog",
				IsDone: false,
			},
			3: {
				ID:     3,
				Text:   "learn dataloaders",
				IsDone: true,
			},
		},
	}
}

func (d db) get(id int) Todo {
	fmt.Println("single - calling database")

	return d.data[id]
}

func (d db) getBatch(ids []int) []Todo {
	fmt.Println("batch - calling database", len(ids))

	var todos []Todo

	for _, id := range ids {
		todos = append(todos, d.data[id])
	}

	return todos
}
