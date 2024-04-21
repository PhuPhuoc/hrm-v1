package common

import "strings"

func GetQueryByFilterObject(filter map[string]interface{}, fields, table string, list_field_filter []string) string {
	var query strings.Builder
	query.WriteString(`select ` + fields + ` from` + table)
	query.WriteString(` where `)
	if len(filter) > 0 || filter != nil {
		// for _, key := range list_field_filter {

		// }
	}
	query.WriteString(`delelted_time is null`)

	return query.String()
}
