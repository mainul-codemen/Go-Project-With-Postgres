package handler

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

type (
	templateData struct {
		Name string
		Age  int
	}
)

func (s *Server) getHome(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"title": func(name string) string {
			return strings.Title(name)
		},
	}

	tmp, _ := template.New("home.html").Funcs(funcMap).ParseFiles("./assets/templates/home.html")

	tmpData := templateData{
		Name: "rahim karim",
		Age:  23,
	}

	err := tmp.Execute(w, tmpData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
