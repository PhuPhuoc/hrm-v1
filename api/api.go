package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/PhuPhuoc/hrm-v1/middleware"
	acc_sv "github.com/PhuPhuoc/hrm-v1/service/account_services"
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
	acc_store := acc_sv.NewAccountStore(sv.db)
	acc_ctl := acc_sv.NewAccooutController(acc_store)
	acc_ctl.RegisterAccountRouter(subrouter)
	/* === */

	log.Printf("<<>> database has connected successfully & the server is listening at port %v", sv.address)

	return http.ListenAndServe(sv.address, router)
}
