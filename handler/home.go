package handler

import (
	"log"
	"net/http"
	"text/template"
)

func (s *Server) getHome(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.New("home.html").ParseFiles("./assets/templates/home.html")

	err := tmp.Execute(w, nil)

	if err != nil {
		log.Println("Error Executing template : ", err)
		return
	}
}
