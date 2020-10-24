package todos

import (
	"context"
	"strconv"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/patrickmn/go-cache"
)

type Repo struct {
	db     db
	loader *dataloader.Loader
}

func NewRepo() Repo {
	repo := Repo{
		db: NewDB(),
	}

	c := cache.New(15*time.Minute, 15*time.Minute)
	cache := Cache{c}
	// cache := dataloader.NoCache{}
	loader := dataloader.NewBatchedLoader(repo.batchFunc, dataloader.WithCache(&cache))

	repo.loader = loader

	return repo
}

func (r Repo) OldGet(id int) Todo {
	return r.db.get(id)
}

func (r Repo) NewGet(id int) Todo {
	result, err := r.loader.Load(context.TODO(), dataloader.StringKey(strconv.Itoa(id)))()
	if err != nil {
		panic(err)
	}

	todo, ok := result.(Todo)
	if !ok {
		panic("failed converting")
	}

	return todo
}

func (r Repo) batchFunc(_ context.Context, keys dataloader.Keys) []*dataloader.Result {
	var todoIDs []int
	for _, key := range keys {
		todoID, err := strconv.Atoi(key.String())
		if err != nil {
			panic(err)
		}

		todoIDs = append(todoIDs, todoID)
	}

	todos := r.db.getBatch(todoIDs)

	var results []*dataloader.Result
	for _, todo := range todos {
		results = append(results, &dataloader.Result{
			Data:  todo,
			Error: nil,
		})
	}

	return results
}
