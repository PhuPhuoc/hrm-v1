package accountservices

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/PhuPhuoc/hrm-v1/common"
	"github.com/PhuPhuoc/hrm-v1/model/account"
	accountquery "github.com/PhuPhuoc/hrm-v1/query/account_query"
)

type accountStore struct {
	db *sql.DB
}

func NewAccountStore(db *sql.DB) *accountStore {
	return &accountStore{
		db: db,
	}
}

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

func (store *accountStore) LoginAccount(email, pwd string) (*account.Account, error) {
	return &account.Account{}, nil
}

func scanRowIntoObjectAccout(row *sql.Row) (*account.Account, error) {
	obj := new(account.Account)

	err := row.Scan(&obj.Id, &obj.FirstName, &obj.LastName, &obj.Email, &obj.PasswordHash, &obj.CreatedTime)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
