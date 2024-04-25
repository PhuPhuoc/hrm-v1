package account

import (
	"github.com/PhuPhuoc/hrm-v1/common"
)

type AccountStore interface {
	CreateAccount(*Account_Register) error
	CheckAccountExistByEmail(string) (bool, error)
	LoginAccount(email, pwd string) (*Account, error)
	GetAllAccount(filter map[string]interface{}, page_num, page_total string) ([]Account, common.Pagination, error)
}
