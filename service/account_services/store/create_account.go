package store

import (
	"fmt"
	"time"

	"github.com/PhuPhuoc/hrm-v1/common"
	"github.com/PhuPhuoc/hrm-v1/model/account"
	accountquery "github.com/PhuPhuoc/hrm-v1/query/account_query"
)

func createDateTimeCurrentFormated() string {

	currentTime := time.Now()
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	currentTimeInLocal := currentTime.In(loc)
	formattedTime := currentTimeInLocal.Format("2006-01-02 15:04:05")

	return formattedTime
}

func (store *accountStore) CreateAccount(acc *account.Account_Register) error {
	pwd_hash := common.GenerateHash(acc.Password)
	query := accountquery.QueryCreateNewAccount()
	result, err := store.db.Exec(query, acc.FirstName, acc.LastName, acc.Email, pwd_hash, acc.AccountRole, createDateTimeCurrentFormated())
	if err != nil {
		return fmt.Errorf("error when CreateAccount in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when CreateAccount in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil
	} else {
		return fmt.Errorf("error when CreateAccount in store (No user created): %v", err)
	}
}
