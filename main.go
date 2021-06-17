package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", getHome).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

func getHome(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.New("home.html").ParseFiles("./assets/templates/home.html")

	err := tmp.Execute(w, nil)

	if err != nil {
		log.Println("Error Executing template : ", err)
		return
	}

}
