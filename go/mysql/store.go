package car

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type query int

const (
	createCarQuery = iota + 1
	readCarQuery
	readAllCarsQuery
	updateCarQuery
	deleteCarQuery
)

var queries = map[query]string{
	createCarQuery: `
	INSERT INTO car(make, model, year, is_new)
	VALUES(?, ?, ?, ?)
	`,
	readCarQuery: `
	SELECT id, make, model, year, is_new FROM car
	WHERE id = ?
	`,
	readAllCarsQuery: `
	SELECT id, make, model, year, is_new FROM car
	`,
	updateCarQuery: `
	UPDATE car
	SET
		make = ?,
		model = ?,
		year = ?,
		is_new = ?
	WHERE id = ?
	`,
	deleteCarQuery: `
	DELETE FROM car WHERE id = ?
	`,
}

// Store reads from mysql
type Store struct {
	db         *sql.DB
	statements map[query]*sql.Stmt
}

// New creates new store with prepared statements
func New() (Store, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mysql-go-testing")
	if err != nil {
		return Store{}, errors.Wrap(err, "failed to open connection to db")
	}

	statements := make(map[query]*sql.Stmt)

	// Prepare statements
	// Info: http://go-database-sql.org/prepared.html
	for key, query := range queries {
		statement, err := db.Prepare(query)
		if err != nil {
			return Store{}, errors.Wrap(err, "failed to prepare statement")
		}

		statements[key] = statement
	}

	return Store{
		db:         db,
		statements: statements,
	}, nil
}

// Ping tries to ping the database
func (store Store) Ping() (string, error) {
	err := store.db.Ping()
	if err != nil {
		return "", errors.Wrap(err, "failed to ping")
	}

	return "pong", nil
}

// Add adds a new car in the database
func (store Store) Add(car Car) (Car, error) {
	stmt := store.statements[createCarQuery]
	defer stmt.Close()

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
func (store Store) Get(id int) (Car, error) {
	stmt := store.statements[readCarQuery]
	defer stmt.Close()

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
func (store Store) GetAll() ([]Car, error) {
	stmt := store.statements[readAllCarsQuery]
	defer stmt.Close()

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
func (store Store) Update(id int, car Car) (Car, error) {
	stmt := store.statements[updateCarQuery]
	defer stmt.Close()

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
func (store Store) Delete(id int) error {
	stmt := store.statements[deleteCarQuery]
	defer stmt.Close()

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
