package oauth

import "github.com/PhuPhuoc/hrm-v1/model/account"

type OauthStore interface {
	LoginWithEmailByOauth(string) (*account.Account, error)
}
