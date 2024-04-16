package account

import "time"

type Account struct {
	Id           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	AccountRole  string    `json:"account_role"`
	CreatedTime  time.Time `json:"-"`
}
