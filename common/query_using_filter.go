package common

import (
	"strconv"
	"strings"
)

func isEmptyMap(m map[string]interface{}) bool {
	// Kiểm tra xem slice của map có độ dài khác không
	// Nếu không, map được coi là trống
	for _, v := range m {
		if len(v.([]interface{})) > 0 {
			return false
		}
	}
	return true
}

func GetQueryByFilterObject(filter map[string]interface{}, fields, table string, list_field_filter []string, page_num, page_total string) string {
	var query strings.Builder
	query.WriteString(`select ` + fields + ` from ` + table)
	query.WriteString(` where `)
	if !isEmptyMap(filter) {
		flag := false
		for _, key := range list_field_filter {
			value_filter, exist := filter[key]
			if exist {
				if flag {
					query.WriteString(` and `)
				}
				if strings.Contains(key, "_from") {
					result := strings.Replace(key, "_from", "", -1)
					query.WriteString(result + ` >= ` + `'` + value_filter.(string) + `'`)
				} else if strings.Contains(key, "_to") {
					result := strings.Replace(key, "_to", "", -1)
					query.WriteString(result + ` <= ` + `'` + value_filter.(string) + `'`)
				} else {
					query.WriteString(key + ` ilike ` + `'%` + value_filter.(string) + `%'`)
				}
				flag = true
			}
		}
		query.WriteString(` and `)
	}
	query.WriteString(`deleted_time is null `)
	query.WriteString(`order by created_time desc`)
	query.WriteString(` limit ` + page_total)
	page, _ := strconv.Atoi(page_num)
	total, _ := strconv.Atoi(page_total)
	query.WriteString(` offset ` + strconv.Itoa((page-1)*total))
	return query.String()
}
