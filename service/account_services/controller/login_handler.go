package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/PhuPhuoc/hrm-v1/common"
	"github.com/PhuPhuoc/hrm-v1/middleware"
	"github.com/PhuPhuoc/hrm-v1/model/account"
)

func (c *accountController) handleAccountLogin(rw http.ResponseWriter, req *http.Request) {
	payload := new(account.RequestLogin)
	acc := new(account.Account)

	var body_data bytes.Buffer
	if _, err_read_body := body_data.ReadFrom(req.Body); err_read_body != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err_read_body))
		return
	}

	json.Unmarshal(body_data.Bytes(), payload)

	exist, err_check_exist := c.store.CheckAccountExistByEmail(payload.Email)
	if err_check_exist != nil {
		common.WriteJSON(rw, common.ErrorResponse_DB(err_check_exist))
		return
	}

	if !exist {
		common.WriteJSON(rw, common.ErrorResponse_BadRequest("email: "+payload.Email+" does not exist", nil))
		return
	}

	acc, err_login := c.store.LoginAccount(payload.Email, payload.Password)
	if err_login != nil {
		if err_login.Error() == "wrong pwd" {
			common.WriteJSON(rw, common.ErrorResponse_BadRequest("wrong password", nil))
		} else {
			common.WriteJSON(rw, common.ErrorResponse_DB(err_login))
		}
		return
	}

	currentTime := time.Now()
	expirationTime := currentTime.Add(2 * time.Hour)
	expUnix := expirationTime.Unix()

	payload_jwt := map[string]interface{}{
		"id":        acc.Id,
		"user_name": acc.FirstName + " " + acc.LastName,
		"exp_date":  expUnix,
	}

	token, err_create_token := middleware.CreateJWT(payload_jwt)
	if err_create_token != nil {
		common.WriteJSON(rw, common.ErrorResponse_Server(err_create_token))
		return
	}

	data_response := make(map[string]interface{})
	data_response["account_info"] = acc
	data_response["token"] = token

	common.WriteJSON(rw, common.SuccessResponse_Data(data_response))
}
