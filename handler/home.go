package handler

import (
	"log"
	"net/http"
)

func (s *Server) getHomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home : get home page")

	template := s.templates.Lookup("home.html")
	if template == nil {
		errMsg := "Unable to load template"
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	if err := template.Execute(w, nil); err != nil {
		log.Fatal("unable to execute template! : ", err)
		return
	}
}
