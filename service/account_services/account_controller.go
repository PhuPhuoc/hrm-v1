package accountservices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm-v1/common"
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
	r.HandleFunc("/register", c.handleAccountRegister).Methods("POST")
	r.HandleFunc("/login", c.handleAccountLogin).Methods("POST")

}

func (c *accountController) handleAccountRegister(rw http.ResponseWriter, req *http.Request) {
	payload := new(account.Account_Register)
	registration := account.Account_Register{}

	// read req.body into bodyData to check valid
	// => when server read req.body 1 time, it will clear the data in req.body
	// => so it need to save a data into a variable to handle other later
	var bodyData bytes.Buffer
	_, err := bodyData.ReadFrom(req.Body)
	if err != nil {
		common.WriteJSON(rw, common.ErrorResponse_Server(err))
		return
	}

	if err_payload := common.ValidateRequestParam(bodyData.Bytes(), registration); err_payload != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err_payload))
		return
	}

	// todo: decode json into payload
	json.Unmarshal(bodyData.Bytes(), payload)

	// todo: check accout exist or not
	flag_exist, err_check_exist := c.store.CheckAccountExistByEmail(payload.Email)
	if err_check_exist != nil {
		common.WriteJSON(rw, common.ErrorResponse_Server(err_check_exist))
		return
	}
	if flag_exist {
		common.WriteJSON(rw, common.ErrorResponse_BadRequest(
			fmt.Sprintf("email: %v already exists", payload.Email),
			fmt.Errorf("email exist in db")))
		return
	}
	// todo: create new accout
	if err_create_acc := c.store.CreateAccount(payload); err_create_acc != nil {
		common.WriteJSON(rw, common.ErrorResponse_Server(err_create_acc))
		return
	}
	common.WriteJSON(rw, common.SuccessResponse_Message("New account registration successful"))
}

func (c *accountController) handleAccountLogin(rw http.ResponseWriter, req *http.Request) {

}
