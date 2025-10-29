package api

import (
	"database/sql"
	"github.com/R-Abinav/GoCommerce/services/user"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

type APIServer struct {
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer{
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error{
	router := mux.NewRouter();
	subRouter := router.PathPrefix("/api/v1").Subrouter();

	userHandler := user.NewHandler();
	userHandler.RegisterRoutes(subRouter);

	log.Println("Server listening on", s.addr);

	return http.ListenAndServe(s.addr, router);

}