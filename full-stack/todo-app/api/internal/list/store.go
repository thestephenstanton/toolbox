package list

import (
	"database/sql"
	"fmt"

	"example.com/todo-api/internal/uid"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type query int

const (
	createListQuery = iota + 1
	readListQuery
	readAllListsQuery
	updateListQuery
	deleteListQuery
)

var queries = map[query]string{
	createListQuery: `
	INSERT INTO list(uid, name)
	VALUES(?, ?)
	`,
	readListQuery: `
	SELECT uid, name 
	FROM list
	WHERE uid = ?
	`,
	readAllListsQuery: `
	SELECT uid, name 
	FROM list
	`,
	updateListQuery: `
	UPDATE list
	SET
		name = ?
	WHERE uid = ?
	`,
	deleteListQuery: `
	DELETE FROM list 
	WHERE uid = ?
	`,
}

// Store reads from mysql
type Store struct {
	db         *sql.DB
	statements map[query]*sql.Stmt
}

// NewStore creates new store with prepared statements
func NewStore() (Store, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mysql-full-stack-todo")
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

// Add adds a new list in the database
func (store Store) Add(list List) (List, error) {
	stmt := store.statements[createListQuery]

	list.UID = uid.NewUID("lst")

	result, err := stmt.Exec(list.UID, list.Name)
	if err != nil {
		return List{}, errors.Wrap(err, "failed to execute prepared statement")
	}

	count, err := result.RowsAffected()
	if err != nil {
		return List{}, errors.Wrap(err, "failed to get rows affected")
	}
	if count == 0 {
		return List{}, errors.New("no rows were added")
	}

	return list, nil
}

// Get gets a list by uid
func (store Store) Get(uid string) (List, error) {
	stmt := store.statements[readListQuery]

	row := stmt.QueryRow(uid)

	var list List
	err := row.Scan(
		&list.UID,
		&list.Name,
	)
	if err != nil {
		return List{}, errors.Wrap(err, fmt.Sprintf("failed to read list with uid %s", uid))
	}

	return list, nil
}

// GetAll gets all lists
func (store Store) GetAll() ([]List, error) {
	stmt := store.statements[readAllListsQuery]

	rows, err := stmt.Query()
	if err != nil {
		return []List{}, errors.Wrap(err, "failed to query all rows")
	}

	var lists []List

	for rows.Next() {
		var list List
		err := rows.Scan(
			&list.UID,
			&list.Name,
		)
		if err != nil {
			return []List{}, errors.Wrap(err, "failed to read list row")
		}

		lists = append(lists, list)
	}

	return lists, nil
}

// Update updates a list given a uid
func (store Store) Update(uid string, list List) (List, error) {
	stmt := store.statements[updateListQuery]

	result, err := stmt.Exec(list.Name, uid)
	if err != nil {
		return List{}, errors.Wrap(err, fmt.Sprintf("failed to update list with uid %s", uid))
	}

	count, err := result.RowsAffected()
	if err != nil {
		return List{}, errors.Wrap(err, fmt.Sprintf("failed to get rows affected"))
	}
	if count == 0 {
		return List{}, fmt.Errorf("no rows were updated with uid %s", uid)
	}

	list.UID = uid

	return list, nil
}

// Delete deletes a list given a uid
func (store Store) Delete(uid string) error {
	stmt := store.statements[deleteListQuery]

	result, err := stmt.Exec(uid)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to delete list with uid %s", uid))
	}

	count, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to get rows affected"))
	}
	if count == 0 {
		return fmt.Errorf("no rows deleted with id %s", uid)
	}

	return nil
}
