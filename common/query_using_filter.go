package common

import (
	"strconv"
	"strings"
)

func isEmptyMap(m map[string]interface{}) bool {
	isEmpty := true

	for range m {
		isEmpty = false
		break
	}
	return isEmpty
}

func GetQueryByFilterObject(filter map[string]interface{}, fields, table string, list_field_filter []string, page_num, page_total string) string {
	var query strings.Builder
	where := CreateConditionClause(filter, list_field_filter)
	part_with := CreateWithClause(where, table)
	part_main := CreateMainClause(where, fields, table)
	part_order := CreateOrderClause()
	part_pagin := CreatePaginationClause(page_num, page_total)
	query.WriteString(part_with + ` ` + part_main + ` ` + part_order + ` ` + part_pagin)
	return query.String()
}

func CreateMainClause(where, fields, table string) string {
	var query strings.Builder
	query.WriteString(`select ` + fields + `, cte.total_record` + ` from ` + table)
	query.WriteString(` cross join cte `)
	query.WriteString(where)
	return query.String()
}

func CreateWithClause(where, table string) string {
	var query strings.Builder
	query.WriteString(`with cte as ( select count(*) as total_record from ` + table + ` ` + where + ` )`)
	return query.String()
}

func CreateConditionClause(filter map[string]interface{}, list_field_filter []string) string {
	var query strings.Builder
	query.WriteString(`where `)
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
				} else if strings.Contains(strings.ToLower(key), "id") {
					if floatValue, ok := value_filter.(float64); ok {
						query.WriteString(`cast(` + key + ` as text)` + ` ilike ` + `'%` + strconv.FormatFloat(floatValue, 'f', -1, 64) + `%'`)
					} else {
						query.WriteString(`cast(` + key + ` as text)` + ` ilike ` + `'%` + value_filter.(string) + `%'`)
					}
				} else if strings.Contains(strings.ToLower(key), "_role") {
					query.WriteString(`cast(` + key + ` as text)` + ` ilike ` + `'%` + value_filter.(string) + `%'`)
				} else {
					query.WriteString(key + ` ilike ` + `'%` + value_filter.(string) + `%'`)
				}
				flag = true
			}
		}
		query.WriteString(` and `)
	}
	query.WriteString(`deleted_time is null`)
	return query.String()
}

func CreateOrderClause() string {
	var query strings.Builder
	query.WriteString(`order by created_time desc`)
	return query.String()
}

func CreatePaginationClause(page_num, page_total string) string {
	var query strings.Builder
	query.WriteString(`limit ` + page_total)
	page, _ := strconv.Atoi(page_num)
	total, _ := strconv.Atoi(page_total)
	query.WriteString(` offset ` + strconv.Itoa((page-1)*total))
	return query.String()
}
