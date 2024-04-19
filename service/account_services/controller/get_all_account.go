package controller

import (
	"net/http"

	"github.com/PhuPhuoc/hrm-v1/common"
)

//	@Summary		get all account
//	@Description	role admin: get all account
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	common.success_response	"Get all account successful"
//	@Failure		400	{object}	common.error_response	"Get all account failure"
//	@Router			/api/v1/account [get]
func (c *accountController) handleGetAllAccount(rw http.ResponseWriter, req *http.Request) {
	common.WriteJSON(rw, common.SuccessResponse_Message("Get into get all account"))
}
