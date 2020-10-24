package transactions

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type query int

const (
	createParentQuery = iota + 1
	createChildQuery
)

var queries = map[query]string{
	createParentQuery: `
	INSERT INTO parent(name)
	VALUES(?)
	`,
	createChildQuery: `
	INSERT INTO child(name, parent_id)
	VALUES(?, ?)
	`,
}

// Store reads from mysql
type Store struct {
	db         *sql.DB
	statements map[query]*sql.Stmt
}

// New creates new store with prepared statements
func New() (Store, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/toolbox-go-transactions")
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

func (store Store) AddMultipleParents(parents []Parent) error {
	fmt.Printf("%v\n\n\n\n\n\n", parents)
	stmt := store.statements[createParentQuery]
	for _, parent := range parents {
		result, err := stmt.Exec(parent.Name)
		if err != nil {
			panic(err)
		}

		count, err := result.RowsAffected()
		if err != nil {
			panic(err)
		}

		if count == 0 {
			panic("no rows affected")
		}
	}

	return nil
}

func (store Store) AddParentAndChild(parent Parent, child Child) error {
	tx, err := store.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	defer tx.Rollback()

	stmt := tx.Stmt(store.statements[createParentQuery])

	result, err := stmt.Exec(parent.Name)
	if err != nil {
		// err = tx.Rollback()
		// if err != nil {
		// 	panic(err.Error())
		// }

		return errors.Wrap(err, "failed to create parent")
	}

	parentID, err := result.LastInsertId()
	if err != nil {
		// err = tx.Rollback()
		// if err != nil {
		// 	panic(err.Error())
		// }

		return errors.Wrap(err, "failed to get last instert id")
	}

	child.ParentID = int(parentID)

	stmt = tx.Stmt(store.statements[createChildQuery])

	result, err = stmt.Exec(child.Name, child.ParentID)
	if err != nil {
		// err = tx.Rollback()
		// if err != nil {
		// 	panic(err.Error())
		// }

		return errors.Wrap(err, "failed to create child")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// err = tx.Rollback()
		// if err != nil {
		// 	panic(err.Error())
		// }

		return errors.Wrap(err, "failed to get rows affected")
	}

	if rowsAffected == 1 {
		// err = tx.Rollback()
		// if err != nil {
		// 	panic(err.Error())
		// }

		return errors.New("no rows were affected")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}
