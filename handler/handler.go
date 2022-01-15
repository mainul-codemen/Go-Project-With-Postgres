package handler

import (
	"Go-Project-With-Postgres/storage/postgres"
	"html/template"
	"io/fs"
	"net/http"
	"path"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/benbjohnson/hashfs"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	homePath    = "/"
	aboutPath   = "/about"
	teamPath    = "/our-team"
	contactPath = "/contact-us"
	signin      = "/signin"
	signup      = "/signup"
)

type TemplateData struct {
	Env       string
	CSRFField template.HTML
}

type Server struct {
	templates *template.Template
	env       string
	logger    *logrus.Entry
	assets    fs.FS
	assetFS   *hashfs.FS
	decoder   *schema.Decoder
	// Mainul Added
	store *postgres.Storage
}

func New(env string,
	config *viper.Viper,
	logger *logrus.Entry,
	assets fs.FS,
	decoder *schema.Decoder) (*mux.Router, error) {

	s := &Server{
		env:     env,
		logger:  logger,
		assets:  assets,
		assetFS: hashfs.NewFS(assets),
		decoder: decoder,
	}

	if err := s.parseTemplates(); err != nil {
		return nil, err
	}

	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", cacheStaticFiles(http.FileServer(http.FS(s.assetFS)))))
	/* -------------------------------------------------------- ROUTES/HANDLERS----------------------------------------------------*/
	r.HandleFunc(homePath, s.getHomeHandler).Methods("GET")
	r.HandleFunc(signup, s.getSignupPage).Methods("GET")
	r.HandleFunc(signup, s.postSignupUser).Methods("POST")
	r.HandleFunc("/event-type", s.createEventType).Methods("GET")
	r.HandleFunc("/event-type", s.saveEventType).Methods("POST")
	r.HandleFunc("/event-type-list", s.getEventType).Methods("GET")

	r.NotFoundHandler = s.getErrorHandler()
	return r, nil
}

func cacheStaticFiles(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if asset is hashed extend cache to 180 days
		e := `"4FROTHS24N"`
		w.Header().Set("Etag", e)
		w.Header().Set("Cache-Control", "max-age=15552000")
		if match := r.Header.Get("If-None-Match"); match != "" {
			if strings.Contains(match, e) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func (s *Server) lookupTemplate(name string) *template.Template {
	if s.env == "development" {
		if err := s.parseTemplates(); err != nil {
			s.logger.WithError(err).Error("template reload")
			return nil
		}
	}
	return s.templates.Lookup(name)
}

func (s *Server) getErrorHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := s.doTemplate(w, r, "error.html", http.StatusTemporaryRedirect); err != nil {
			s.logger.WithError(err).Error("unable to load error template")
		}
	})
}

func (s *Server) templateData(r *http.Request) TemplateData {
	return TemplateData{
		Env:       s.env,
		CSRFField: csrf.TemplateField(r),
	}
}

func (s *Server) doTemplate(w http.ResponseWriter, r *http.Request, name string, status int) error {
	template := s.lookupTemplate(name)
	if template == nil || isPartialTemplate(name) {
		template, status = s.templates.Lookup("error.html"), http.StatusNotFound
	}

	w.WriteHeader(status)
	return template.Execute(w, s.templateData(r))
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-forwarded-for")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func isPartialTemplate(name string) bool {
	return strings.HasSuffix(name, ".part.html")
}

func (s *Server) parseTemplates() error {
	templates := template.New("templates").Funcs(template.FuncMap{
		"assetHash": func(n string) string {
			return path.Join("/", s.assetFS.HashName(strings.TrimPrefix(path.Clean(n), "/")))
		},
	}).Funcs(sprig.FuncMap())

	tmpl, err := templates.ParseFS(s.assets, "templates/*/*.html")
	if err != nil {
		return err
	}
	s.templates = tmpl
	return nil
}
