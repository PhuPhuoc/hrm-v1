package account

// RequestLogin
type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
