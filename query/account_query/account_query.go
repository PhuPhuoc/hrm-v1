package accountquery

var table = `accounts`
var account_field = ` ` + `id, first_name, last_name, email, password_hash, account_role, created_time` + ` `
var table_name = ` ` + table + ` `

func QueryCheckAccountExistByEmail() string {
	return `select id from` + table_name + `where email = $1`
}

func QueryCreateNewAccount() string {
	return `insert into` + table_name + `(first_name, last_name, email, password_hash, account_role, created_time) values ($1,$2,$3,$4,$5,$6)`
}

func QueryLogin() string {
	return `select` + account_field + `from` + table_name + ` where email = $1 and deleted_time is null`
}
