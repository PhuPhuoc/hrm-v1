package oauthgoogleservices

import (
	"context"
	"net/http"

	"github.com/PhuPhuoc/hrm-v1/common"
	"github.com/PhuPhuoc/hrm-v1/model/oauth"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

type OauthApp struct {
	conf  oauth2.Config
	store oauth.OauthStore
}

func InitOauth(config oauth2.Config, store oauth.OauthStore) *OauthApp {
	return &OauthApp{
		conf:  config,
		store: store,
	}
}

func (a *OauthApp) RegisterAuthRouter(r *mux.Router) {
	auth_router := r.PathPrefix("/auth").Subrouter()
	auth_router.HandleFunc("/oauth", a.HandleOauth).Methods("GET")
	auth_router.HandleFunc("/callback", a.HandleCallback).Methods("GET")
	auth_router.HandleFunc("/login-google-account", a.HandleGetUserInfoToLogin).Methods("POST")
}

func (a *OauthApp) HandleOauth(w http.ResponseWriter, r *http.Request) {
	url := a.conf.AuthCodeURL("HRM_version1", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *OauthApp) HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// Exchanging the code for an access token
	t, err := a.conf.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	common.WriteJSON(w, t)
}
