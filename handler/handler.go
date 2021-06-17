package handler

import "github.com/gorilla/mux"

type (
	Server struct {
	}
)

func NewServer() (*mux.Router, error) {

	s := &Server{}

	r := mux.NewRouter()

	r.HandleFunc("/", s.getHome).Methods("GET")
	return r, nil
}
