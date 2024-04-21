package account

import "time"

type AccountFilter struct {
	Id               int       `json:"id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Email            string    `json:"email"`
	AccountRole      string    `json:"account_role" enums:"HR,ADMIN"`
	CreatedTime_From time.Time `json:"created_time_from"`
	CreatedTime_To   time.Time `json:"created_time_to"`
}
