package account

type AccountStore interface {
	CreateAccount(*Account_Register) error
	CheckAccountExistByEmail(string) (bool, error)
	LoginAccount(email, pwd string) (*Account, error)
}
