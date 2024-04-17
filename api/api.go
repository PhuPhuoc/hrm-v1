package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/PhuPhuoc/hrm-v1/middleware"
	acc_ctl "github.com/PhuPhuoc/hrm-v1/service/account_services/controller"
	acc_store "github.com/PhuPhuoc/hrm-v1/service/account_services/store"
	"github.com/gorilla/mux"
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

	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the server: HRM-v1")
	})

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	/* account */
	account_store := acc_store.NewAccountStore(sv.db)
	account_controller := acc_ctl.NewAccooutController(account_store)
	account_controller.RegisterAccountRouter(subrouter)

	log.Printf("<<>> database has connected successfully & the server is listening at port %v", sv.address)

	return http.ListenAndServe(sv.address, router)
}
