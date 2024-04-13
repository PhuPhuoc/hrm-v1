package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the server: HRM-v1")
	})

	//subrouter := router.PathPrefix("api/v1").Subrouter()

	log.Printf("<<>> database has connected successfully & the server is listening at port %v", sv.address)

	return http.ListenAndServe(sv.address, router)
}
