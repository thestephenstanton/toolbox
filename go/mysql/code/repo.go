package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// Repo reads from mysql
type Repo struct {
	db         *sql.DB
	statements map[query]*sql.Stmt
}

// New creates new repo with prepared statements
func New(dataSourceName string) (Repo, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return Repo{}, errors.Wrap(err, "failed to open connection to db")
	}

	statements := map[query]*sql.Stmt{}

	// Prepare statements
	// Info: http://go-database-sql.org/prepared.html
	for key, query := range queries {
		statement, err := db.Prepare(query)
		if err != nil {
			return Repo{}, errors.Wrap(err, "failed to prepare statement")
		}

		statements[key] = statement
	}

	repo := Repo{
		db:         db,
		statements: statements,
	}

	_, err = repo.Ping()
	if err != nil {
		return Repo{}, errors.Wrap(err, "failed to ping database")
	}

	return repo, nil
}

// Ping tries to ping the database
func (repo Repo) Ping() (string, error) {
	err := repo.db.Ping()
	if err != nil {
		return "", errors.Wrap(err, "failed to ping")
	}

	return "pong", nil
}

// Add adds a new car in the database
func (repo Repo) Add(car Car) (Car, error) {
	stmt := repo.statements[createCarQuery]

	result, err := stmt.Exec(car.Make, car.Model, car.Year, car.IsNew)
	if err != nil {
		return Car{}, errors.Wrap(err, "failed to execute prepared statement")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Car{}, errors.Wrap(err, "failed to get last insert id")
	}

	car.ID = int(id)

	return car, nil
}

// Get gets a car by id
func (repo Repo) Get(id int) (Car, error) {
	stmt := repo.statements[readCarQuery]

	row := stmt.QueryRow(id)

	var car Car
	err := row.Scan(
		&car.ID,
		&car.Make,
		&car.Model,
		&car.Year,
		&car.IsNew,
	)
	if err != nil {
		return Car{}, errors.Wrap(err, fmt.Sprintf("failed to read car with id %d", id))
	}

	return car, nil
}

// GetAll gets all cars
func (repo Repo) GetAll() ([]Car, error) {
	stmt := repo.statements[readAllCarsQuery]

	rows, err := stmt.Query()
	if err != nil {
		return []Car{}, errors.Wrap(err, "failed to query all rows")
	}

	var cars []Car

	for rows.Next() {
		var car Car
		err := rows.Scan(
			&car.ID,
			&car.Make,
			&car.Model,
			&car.Year,
			&car.IsNew,
		)
		if err != nil {
			return []Car{}, errors.Wrap(err, "failed to read car row")
		}

		cars = append(cars, car)
	}

	return cars, nil
}

// Update updates a car given an id
func (repo Repo) Update(id int, car Car) (Car, error) {
	stmt := repo.statements[updateCarQuery]

	result, err := stmt.Exec(car.Make, car.Model, car.Year, car.IsNew, id)
	if err != nil {
		return Car{}, errors.Wrap(err, fmt.Sprintf("failed to update car with id %d", id))
	}

	count, err := result.RowsAffected()
	if err != nil {
		return Car{}, errors.Wrap(err, fmt.Sprintf("failed to get rows affected"))
	}
	if count == 0 {
		return Car{}, fmt.Errorf("no rows were updated with id %d", id)
	}

	car.ID = id

	return car, nil
}

// Delete deletes a car given an id
func (repo Repo) Delete(id int) error {
	stmt := repo.statements[deleteCarQuery]

	result, err := stmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to delete car with id %d", id))
	}

	count, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to get rows affected"))
	}
	if count == 0 {
		return fmt.Errorf("no rows deleted with id %d", id)
	}

	return nil
}
