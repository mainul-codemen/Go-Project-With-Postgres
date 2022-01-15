package handler

import (
	"Go-Project-With-Postgres/storage"
	"html/template"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/csrf"
)

type UserFormData struct {
	CSRFField  template.HTML
	Form       storage.User
	FormErrors map[string]string
}

func (s *Server) getSignupPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Method : Create user called.")
	data := UserFormData{
		CSRFField: csrf.TemplateField(r),
	}
	s.loadUserTemplate(w, r, data)

}

func (s *Server) postSignupUser(w http.ResponseWriter, r *http.Request) {
	log.Panicln("Method : Post sign up called")
	ParseFormData(r)
	var creds storage.User
	if err := s.decoder.Decode(&creds, r.PostForm); err != nil {
		log.Fatalln("Decoding error")
	}
	// validation
	if err := creds.Validate(); err != nil {
		vErrs := map[string]string{}
		if e, ok := err.(validation.Errors); ok {
			if len(e) > 0 {
				for key, value := range e {
					vErrs[key] = value.Error()
				}
			}
		}
		data := UserFormData{
			CSRFField:  csrf.TemplateField(r),
			Form:       creds,
			FormErrors: vErrs,
		}
		s.loadUserTemplate(w, r, data)
		return
	}

	pass := creds.Password
	hashed, _ := HashAndSalt(pass)
	creds.Password = hashed
	_, err := s.store.CreateUser(creds)
	UnableToInsertData(err)
	http.Redirect(w, r, "/event", http.StatusSeeOther)
}

func (s *Server) loadUserTemplate(w http.ResponseWriter, r *http.Request, form UserFormData) {
	tmpl := s.templates.Lookup("user-form.html")
	UnableToFindHtmlTemplate(tmpl)
	err := tmpl.Execute(w, form)
	ExcutionTemplateError(err)
}
