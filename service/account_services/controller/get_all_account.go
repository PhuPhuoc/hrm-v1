package controller

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/PhuPhuoc/hrm-v1/common"
)

// @Summary		Get all accounts
// @Description	Get all accounts. Requires admin role.
// @Tags			Account
// @Accept			json
// @Produce		json
// @Param			request	body		account.AccountFilter	true	"Get all accounts request"
// @Success		200		{object}	common.success_response	"Get all accounts successful"
// @Failure		400		{object}	common.error_response	"Get all accounts failure"
// @Router			/api/v1/account/get-all [post]
// @Security		ApiKeyAuth
func (c *accountController) handleGetAllAccount(rw http.ResponseWriter, req *http.Request) {
	var bodyData bytes.Buffer
	_, err := bodyData.ReadFrom(req.Body)
	if err != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err))
		return
	}

	var param_json_filter map[string]interface{}
	err = json.Unmarshal(bodyData.Bytes(), &param_json_filter)
	if err != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err))
		return
	}

	common.WriteJSON(rw, common.SuccessResponse_Data(param_json_filter))
}
