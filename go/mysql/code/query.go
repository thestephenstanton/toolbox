package main

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
	INSERT INTO car(make, model, year, is_new, mileage)
	VALUES(?, ?, ?, ?, ?)
	`,
	readCarQuery: `
	SELECT id, make, model, year, is_new, mileage
	FROM car
	WHERE id = ?
	`,
	readAllCarsQuery: `
	SELECT id, make, model, year, is_new, mileage
	FROM car
	`,
	updateCarQuery: `
	UPDATE car
	SET
		make = ?,
		model = ?,
		year = ?,
		is_new = ?,
		mileage = ?
	WHERE id = ?
	`,
	deleteCarQuery: `
	DELETE FROM car 
	WHERE id = ?
	`,
}
