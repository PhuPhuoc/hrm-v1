package store

import (
	"fmt"

	"github.com/PhuPhuoc/hrm-v1/common"
	"github.com/PhuPhuoc/hrm-v1/model/account"
	accountquery "github.com/PhuPhuoc/hrm-v1/query/account_query"
)

func (store *accountStore) LoginAccount(email, pwd string) (*account.Account, error) {
	acc := new(account.Account)
	query := accountquery.QueryLogin()
	if err_query := store.db.QueryRow(query, email).Scan(&acc.Id, &acc.FirstName, &acc.LastName, &acc.Email, &acc.PasswordHash, &acc.AccountRole, &acc.CreatedTime); err_query != nil {
		return acc, err_query
	}

	if !common.CompareHash(pwd, acc.PasswordHash) {
		return acc, fmt.Errorf("wrong pwd")
	}

	return acc, nil
}
