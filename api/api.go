package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/PhuPhuoc/hrm-v1/docs"
	"github.com/PhuPhuoc/hrm-v1/middleware"
	acc_ctl "github.com/PhuPhuoc/hrm-v1/service/account_services/controller"
	acc_store "github.com/PhuPhuoc/hrm-v1/service/account_services/store"
	oauthgoogleservices "github.com/PhuPhuoc/hrm-v1/service/oauth_google_services"
	oauth_store "github.com/PhuPhuoc/hrm-v1/service/oauth_google_services/store"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	address string
	db      *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		address: addr,
		db:      db,
	}
}

func (sv *Server) Run() error {
	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the server: HRM-v1"})
	}).Methods("GET")

	subrouter := router.PathPrefix("/api/v1").Subrouter()
	subrouter.Use(middleware.LoggingMiddleware)

	/* auth-config */
	conf, err_oauth := oauthgoogleservices.NewOauthAppConfig()
	if err_oauth != nil {
		log.Println("error when config oauth: ", err_oauth)
	}
	/* auth */
	oauthStore := oauth_store.NewOauthStore(sv.db)
	oauth := oauthgoogleservices.InitOauth(*conf, oauthStore)
	oauth.RegisterAuthRouter(subrouter)

	/* account */
	account_store := acc_store.NewAccountStore(sv.db)
	account_controller := acc_ctl.NewAccooutController(account_store)
	account_controller.RegisterAccountRouter(subrouter)

	url_api_doc := fmt.Sprintf("http://localhost%v/swagger/index.html", sv.address)
	url_login_google := fmt.Sprintf("http://localhost%v/api/v1/auth/oauth", sv.address)

	log.Printf("\n  ~  Server is listening at port  |%v| \n  ~  API docs -  %v \n  ~  URL login with Google -  %v", sv.address, url_api_doc, url_login_google)

	return http.ListenAndServe(sv.address, router)
}
