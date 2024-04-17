package store

import (
	"database/sql"
)

type accountStore struct {
	db *sql.DB
}

func NewAccountStore(db *sql.DB) *accountStore {
	return &accountStore{
		db: db,
	}
}

// func scanRowIntoObjectAccout(row *sql.Row) (*account.Account, error) {
// 	obj := new(account.Account)

// 	err := row.Scan(&obj.Id, &obj.FirstName, &obj.LastName, &obj.Email, &obj.PasswordHash, &obj.CreatedTime)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return obj, nil
// }
