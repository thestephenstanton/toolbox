package list

// type Store struct {
// 	db      *sql.DB
// 	queries map[query]*sql.Stmt
// }

// func NewStore(db *sql.DB) (Store, error) {
// 	unprepared := map[query]string{
// 		queryAdd: `
// 			INSERT INTO todo(uid, text, list_id)
// 			SELECT ?, ?, l.id
// 			FROM 'list' l
// 			WHERE l.'uid'= 'foo'
// 		`,
// 	}

// 	prepared, err := prepareStatements(db, unprepared)
// 	if err != nil {
// 		return Store{}, errors.Wrap(err, "failed to prepare statements")
// 	}
// }

// func prepareStatements(db *sql.DB, unprepared map[query]string) (map[query]*sql.Stmt, error) {
// 	prepared := map[query]*sql.Stmt{}
// 	for query, statement := range unprepared {
// 		stmt, err := db.Prepare(statement)
// 		if err != nil {
// 			return nil, err
// 		}

// 		prepared[query] = stmt
// 	}

// 	return prepared, nil
// }
