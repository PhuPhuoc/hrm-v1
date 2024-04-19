package account

import (
	"encoding/json"
	"fmt"
)

type EnumAccountRole string

const (
	HR    EnumAccountRole = "HR"
	ADMIN EnumAccountRole = "ADMIN"
)

var validRoles = map[EnumAccountRole]bool{
	HR:    true,
	ADMIN: true,
}

type Account_Register struct {
	FirstName   string          `json:"first_name,omitempty"`
	LastName    string          `json:"last_name,omitempty"`
	Email       string          `json:"email"`
	Password    string          `json:"password"`
	AccountRole EnumAccountRole `json:"account_role" enums:"HR,ADMIN"`
}

func isValidAccountRole(role EnumAccountRole) bool {
	return validRoles[role]
}

func (acc *Account_Register) ConvertBodyDataToModel(bd []byte) error {
	json.Unmarshal(bd, acc)
	if isValidAccountRole(acc.AccountRole) {
		return nil
	}
	return fmt.Errorf("account role must be 'HR' or 'ADMIN'")
}
