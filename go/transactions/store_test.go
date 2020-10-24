package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	store, err := New()
	assert.NoError(t, err)

	result, err := store.Ping()
	assert.NoError(t, err)

	assert.Equal(t, "pong", result)
}

func TestAddMultipleParents(t *testing.T) {
	store, err := New()
	assert.NoError(t, err)

	parents := []Parent{
		Parent{Name: "Fizz"},
		Parent{Name: "Buzz"},
	}

	err = store.AddMultipleParents(parents)
	assert.NoError(t, err)

	t.Error("sdf")

}

func TestAddParentAndChild(t *testing.T) {
	parent := Parent{
		Name: "Joe",
	}

	child := Child{
		Name: "Bob",
	}

	store, err := New()
	assert.NoError(t, err)

	err = store.AddParentAndChild(parent, child)
	assert.NoError(t, err)
}
