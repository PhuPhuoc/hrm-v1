package store

import (
	"database/sql"
	"fmt"

	"github.com/PhuPhuoc/hrm-v1/model/account"
	accountquery "github.com/PhuPhuoc/hrm-v1/query/account_query"
)

type oauthStore struct {
	db *sql.DB
}

func NewOauthStore(db *sql.DB) *oauthStore {
	return &oauthStore{
		db: db,
	}
}

func (store *oauthStore) LoginWithEmailByOauth(email string) (*account.Account, error) {
	acc := new(account.Account)
	query := accountquery.QueryLogin()
	err := store.db.QueryRow(query, email).Scan(&acc.Id, &acc.FirstName, &acc.LastName, &acc.Email, &acc.PasswordHash, &acc.AccountRole, &acc.CreatedTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("email %v does not exist", email)
		}
		return nil, err
	}
	return acc, nil
}
