package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Server struct {
	}
)

func NewServer() (*mux.Router, error) {

	s := &Server{}

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))
	r.HandleFunc("/", s.getHome).Methods("GET")
	return r, nil
}
