package controller

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm-v1/common"
	"github.com/PhuPhuoc/hrm-v1/model/account"
)

//	@Summary		register new account
//	@Description	create new account with user's info
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Param			request	body		account.Account_Register	true	"Register request"
//	@Success		201		{object}	common.success_response		"Created new account successfully"
//	@Failure		400		{object}	common.error_response		"Create failure"
//	@Router			/api/v1/register [post]
func (c *accountController) handleAccountRegister(rw http.ResponseWriter, req *http.Request) {
	payload := new(account.Account_Register)
	registration := account.Account_Register{}
	/*
		todo: read req.body into bodyData to check valid
		? when server read req.body 1 time, it will clear the data in req.body
		? it need to save a data into a variable to handle other later
	*/
	var bodyData bytes.Buffer
	_, err := bodyData.ReadFrom(req.Body)
	if err != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err))
		return
	}

	if err_payload := common.ValidateRequestParam(bodyData.Bytes(), registration); err_payload != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err_payload))
		return
	}

	// todo: decode json into payload
	if err_convert := payload.ConvertBodyDataToModel(bodyData.Bytes()); err_convert != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err_convert))
		return
	}

	// todo: check accout exist or not
	flag_exist, err_check_exist := c.store.CheckAccountExistByEmail(payload.Email)
	if err_check_exist != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err_check_exist))
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
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err_create_acc))
		return
	}
	common.WriteJSON(rw, common.SuccessResponse_Message("New account registration successful"))
}
