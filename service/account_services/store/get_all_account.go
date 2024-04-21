package store

import (
	"fmt"

	"github.com/PhuPhuoc/hrm-v1/common"
	"github.com/PhuPhuoc/hrm-v1/model/account"
)

func (store *accountStore) GetAllAccount(filter map[string]interface{}) ([]account.Account, error) {
	query := common.GetQueryByFilterObject(filter, ACCOUNT_FIELDS, ACCOUNT_TABLE_NAME, ACCOUNT_FILTER)
	fmt.Println(query)
	return nil, nil
}
