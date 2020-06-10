package todo

import (
	"database/sql"
	"fmt"

	"example.com/todo-api/internal/uid"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type query int

const (
	createTodoQuery = iota + 1
	readTodoQuery
	readAllTodosQuery
	updateTodoQuery
	deleteTodoQuery
)

var queries = map[query]string{
	createTodoQuery: `
	INSERT INTO todo(uid, text, list_id)
	SELECT ?, ?, l.id
	FROM list l WHERE l.uid = ?
	`,
	readTodoQuery: `
	SELECT uid, text, is_finished
	FROM todo
	WHERE uid = ?
	`,
	readAllTodosQuery: `
	SELECT uid, text, is_finished
	FROM todo
	`,
	updateTodoQuery: `
	UPDATE todo
	SET
		text = ?,
		is_finished = ?
	WHERE uid = ?
	`,
	deleteTodoQuery: `
	DELETE FROM todo 
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

// Add adds a new todo in the database
func (store Store) Add(todo Todo) (Todo, error) {
	stmt := store.statements[createTodoQuery]

	todo.UID = uid.NewUID("td")

	result, err := stmt.Exec(todo.UID, todo.Text, todo.ListUID)
	if err != nil {
		return Todo{}, errors.Wrap(err, "failed to execute prepared statement")
	}

	count, err := result.RowsAffected()
	if err != nil {
		return Todo{}, errors.Wrap(err, "failed to get rows affected")
	}
	if count == 0 {
		return Todo{}, errors.New("no rows were added")
	}

	return todo, nil
}

// Get gets a todo by uid
func (store Store) Get(uid string) (Todo, error) {
	stmt := store.statements[readTodoQuery]

	row := stmt.QueryRow(uid)

	var todo Todo
	err := row.Scan(
		&todo.UID,
		&todo.Text,
		&todo.IsFinsihed,
		&todo.ListUID,
	)
	if err != nil {
		return Todo{}, errors.Wrap(err, fmt.Sprintf("failed to read todo with uid %d", uid))
	}

	return todo, nil
}

// GetAll gets all todos
func (store Store) GetAll() ([]Todo, error) {
	stmt := store.statements[readAllTodosQuery]

	rows, err := stmt.Query()
	if err != nil {
		return []Todo{}, errors.Wrap(err, "failed to query all rows")
	}

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.UID,
			&todo.Text,
			&todo.IsFinsihed,
			&todo.ListUID,
		)
		if err != nil {
			return []Todo{}, errors.Wrap(err, "failed to read todo row")
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

// Update updates a todo given a uid
func (store Store) Update(uid string, todo Todo) (Todo, error) {
	stmt := store.statements[updateTodoQuery]

	result, err := stmt.Exec(todo.Text, todo.IsFinsihed, uid)
	if err != nil {
		return Todo{}, errors.Wrap(err, fmt.Sprintf("failed to update todo with uid %s", uid))
	}

	count, err := result.RowsAffected()
	if err != nil {
		return Todo{}, errors.Wrap(err, fmt.Sprintf("failed to get rows affected"))
	}
	if count == 0 {
		return Todo{}, fmt.Errorf("no rows were updated with uid %s", uid)
	}

	todo.UID = uid

	return todo, nil
}

// Delete deletes a todo given a uid
func (store Store) Delete(uid string) error {
	stmt := store.statements[deleteTodoQuery]

	result, err := stmt.Exec(uid)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to delete todo with id %s", uid))
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
