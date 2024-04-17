package store

import (
	"fmt"

	accountquery "github.com/PhuPhuoc/hrm-v1/query/account_query"
)

func (store *accountStore) CheckAccountExistByEmail(email string) (bool, error) {
	query_str := accountquery.QueryCheckAccountExistByEmail()
	rows, err := store.db.Query(query_str, email)
	if err != nil {
		return true, fmt.Errorf("error when CheckAccountExistByEmail in store: %v", err)
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}
