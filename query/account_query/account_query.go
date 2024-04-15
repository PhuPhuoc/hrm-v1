package accountquery

// var account_field = `id, first_name, last_name, email, password_hash, created_time`

func QueryCheckAccountExistByEmail() string {
	return `select id from accounts where email = $1`
}

func QueryCreateNewAccount() string {
	return `insert into accounts(first_name, last_name, email, password_hash, created_time) values ($1,$2,$3,$4,$5)`
}
