package controller

import (
	"github.com/PhuPhuoc/hrm-v1/middleware"
	"github.com/PhuPhuoc/hrm-v1/model/account"
	"github.com/gorilla/mux"
)

type accountController struct {
	store account.AccountStore
}

func NewAccooutController(s account.AccountStore) *accountController {
	return &accountController{
		store: s,
	}
}

func (c *accountController) RegisterAccountRouter(r *mux.Router) {

	//r.HandleFunc("/register", c.handleAccountRegister).Methods("POST")
	r.HandleFunc("/login", c.handleAccountLogin).Methods("POST")

	account_management_router := r.PathPrefix("/account").Subrouter()
	account_management_router.Use(middleware.ValidateTokenMiddleware, middleware.AuthorizationMiddleware)
	account_management_router.HandleFunc("", c.handleGetAllAccount).Methods("GET")
	account_management_router.HandleFunc("", c.handleAccountRegister).Methods("POST")

}
