package store

import (
	"database/sql"
)

const ACCOUNT_FIELDS = "id, first_name, last_name, email, account_role, created_time"
const ACCOUNT_TABLE_NAME = "accounts"

var ACCOUNT_FILTER = []string{"id", "first_name", "last_name", "email", "account_role", "created_time_from", "created_time_to"}

type accountStore struct {
	db *sql.DB
}

func NewAccountStore(db *sql.DB) *accountStore {
	return &accountStore{
		db: db,
	}
}
