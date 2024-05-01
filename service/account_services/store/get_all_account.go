package store

import (
	"strconv"

	"github.com/PhuPhuoc/hrm-v1/common"
	"github.com/PhuPhuoc/hrm-v1/model/account"
)

func (store *accountStore) GetAllAccount(filter map[string]interface{}, page_num, page_total string) ([]account.Account, common.Pagination, error) {
	data := []account.Account{}
	pagin := common.Pagination{}
	if err := pagin.Process(page_num, page_total); err != nil {
		return nil, pagin, err
	}

	total_record := 0

	query := common.GetQueryByFilterObject(filter, ACCOUNT_FIELDS, ACCOUNT_TABLE_NAME, ACCOUNT_FILTER, strconv.Itoa(pagin.Current_Page), strconv.Itoa(pagin.Limit))
	//fmt.Println("+ Get All Account Query: ", query)
	rows, err_query := store.db.Query(query)
	if err_query != nil {
		return nil, pagin, err_query
	}
	defer rows.Close()

	for rows.Next() {
		acc := new(account.Account)
		err_scan := rows.Scan(&acc.Id, &acc.FirstName, &acc.LastName, &acc.Email, &acc.AccountRole, &acc.CreatedTime, &total_record)
		if err_scan != nil {
			return nil, pagin, err_query
		}
		data = append(data, *acc)
	}

	pagin.Total_Item = total_record
	per := pagin.Total_Item % pagin.Limit
	if per > 0 {
		pagin.Total_Page = pagin.Total_Item/pagin.Limit + 1
	} else {
		pagin.Total_Page = pagin.Total_Item / pagin.Limit
	}
	return data, pagin, nil
}
