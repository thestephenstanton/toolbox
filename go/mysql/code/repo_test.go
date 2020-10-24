package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	repo, err := New("root:password@tcp(127.0.0.1:3306)/mysql-go-testing")
	assert.NoError(t, err)
	assert.NotNil(t, repo)
}

func TestPing(t *testing.T) {
	store, err := New("root:password@tcp(127.0.0.1:3306)/mysql-go-testing")
	assert.NoError(t, err)

	result, err := store.Ping()
	assert.NoError(t, err)
	assert.Equal(t, "pong", result)
}

func TestAdd(t *testing.T) {
	store, err := New("root:password@tcp(127.0.0.1:3306)/mysql-go-testing")
	assert.NoError(t, err)

	newCar := Car{
		Make:  "Audi",
		Model: "RS7",
		Year:  2020,
		IsNew: true,
	}

	savedCar, err := store.Add(newCar)
	assert.NoError(t, err)

	newCar.ID = savedCar.ID

	assert.Equal(t, savedCar, newCar)
}

func TestGet(t *testing.T) {
	store, err := New("root:password@tcp(127.0.0.1:3306)/mysql-go-testing")
	assert.NoError(t, err)

	oldCar := Car{
		Make:  "Lamborghini",
		Model: "Reventon",
		Year:  2009,
		IsNew: false,
	}

	savedCar, err := store.Add(oldCar)
	assert.NoError(t, err)

	oldCar.ID = savedCar.ID

	foundCar, err := store.Get(oldCar.ID)
	assert.NoError(t, err)

	assert.Equal(t, oldCar, foundCar)
}

func TestGetAll(t *testing.T) {
	store, err := New("root:password@tcp(127.0.0.1:3306)/mysql-go-testing")
	assert.NoError(t, err)

	allCars, err := store.GetAll()
	assert.NoError(t, err)

	assert.NotEmpty(t, allCars)
}

func TestErrors(t *testing.T) {
	testCases := []struct {
		desc          string
		actualErr     error
		shouldError   bool
		expectedError string
	}{
		{
			desc:      "Both nil",
			actualErr: nil,
		},
		{
			desc:          "Regular error",
			actualErr:     errors.New("some random error"),
			shouldError:   true,
			expectedError: "some random error",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			assertCheckError(t, tc.actualErr, tc.expectedError)
			// if tc.shouldError && tc.actualErr == nil {
			// 	assert.Fail(t, "should have errored but didn't")
			// }
			// if tc.actualErr != nil {
			// 	assert.EqualError(t, tc.actualErr, tc.expectedError)
			// }
		})
	}
}

func assertCheckError(t *testing.T, actualErr error, expectedError string) {
	shouldError := len(expectedError) > 0

	if shouldError && actualErr == nil {
		assert.Fail(t, "should have errored but didn't")
	}
	if actualErr != nil {
		assert.EqualError(t, actualErr, expectedError)
	}
}

func TestUpdate(t *testing.T) {
	store, err := New("root:password@tcp(127.0.0.1:3306)/mysql-go-testing")
	assert.NoError(t, err)

	oopsCar := Car{
		Make:  "Rollz-Royce",
		Model: "Wraith",
		Year:  2020,
		IsNew: true,
	}

	carToUpdate, err := store.Add(oopsCar)
	assert.NoError(t, err)

	carToUpdate.Make = "Rolls-Royce"

	updatedCar, err := store.Update(carToUpdate.ID, carToUpdate)
	assert.NoError(t, err)

	assert.Equal(t, carToUpdate, updatedCar)
}

func TestDelete(t *testing.T) {
	store, err := New("root:password@tcp(127.0.0.1:3306)/mysql-go-testing")
	assert.NoError(t, err)

	// Car was a hoax
	shitCar := Car{
		Make:  "Devel",
		Model: "Sixteen",
		Year:  2999,
		IsNew: false,
	}

	shitCar, err = store.Add(shitCar)
	assert.NoError(t, err)

	err = store.Delete(shitCar.ID)
	assert.NoError(t, err)
}
