package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/PhuPhuoc/hrm-v1/docs"
	"github.com/PhuPhuoc/hrm-v1/middleware"
	acc_ctl "github.com/PhuPhuoc/hrm-v1/service/account_services/controller"
	acc_store "github.com/PhuPhuoc/hrm-v1/service/account_services/store"
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

	/* account */
	account_store := acc_store.NewAccountStore(sv.db)
	account_controller := acc_ctl.NewAccooutController(account_store)
	account_controller.RegisterAccountRouter(subrouter)

	log.Printf("<<>> database has connected successfully & the server is listening at port %v", sv.address)

	return http.ListenAndServe(sv.address, router)
}
