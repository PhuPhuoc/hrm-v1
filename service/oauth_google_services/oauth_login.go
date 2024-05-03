package oauthgoogleservices

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PhuPhuoc/hrm-v1/common"
	"github.com/PhuPhuoc/hrm-v1/middleware"
	"golang.org/x/oauth2"
)

// @Summary		Get user information from Google
// @Description	Get user information from Google using OAuth2 token
// @Tags			Auth
// @Produce		json
// @Param			token	body		oauth.TokenResponse										true	"OAuth2 token"
// @Success		200		{object}	common.success_response{data=map[string]interface{}}	"User information retrieved successfully"
// @Failure		400		{object}	common.error_response									"Bad Request"
// @Failure		500		{object}	common.error_response									"Internal Server Error"
// @Router			/api/v1/auth/login-google-account [post]
func (a *OauthApp) HandleGetUserInfoToLogin(rw http.ResponseWriter, r *http.Request) {
	token := &oauth2.Token{}
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err))
		return
	}

	client := a.conf.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err))
		return
	}
	defer resp.Body.Close()

	var user_info map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&user_info)
	if err != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err))
		return
	}

	email, ok := user_info["email"]
	if !ok {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(fmt.Errorf("payload does not have field: email")))
		return
	}

	acc_info, err_query := a.store.LoginWithEmailByOauth(email.(string))
	if err_query != nil {
		common.WriteJSON(rw, common.ErrorResponse_InvalidRequest(err_query))
		return
	}

	currentTime := time.Now()
	expirationTime := currentTime.Add(120 * time.Minute)
	expUnix := expirationTime.Unix()

	payload_jwt := map[string]interface{}{
		"id":        acc_info.Id,
		"role":      acc_info.AccountRole,
		"user_name": acc_info.LastName + " " + acc_info.FirstName,
		"exp_date":  expUnix,
	}

	jwt, err_create_token := middleware.CreateJWT(payload_jwt)
	if err_create_token != nil {
		common.WriteJSON(rw, common.ErrorResponse_Server(err_create_token))
		return
	}

	data_response := make(map[string]interface{})
	data_response["account_info"] = acc_info
	data_response["token"] = jwt

	common.WriteJSON(rw, common.SuccessResponse_Data(data_response))
}
