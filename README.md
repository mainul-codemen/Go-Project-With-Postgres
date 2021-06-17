# Go-Project-With-Postgres

This is a demo project for Learning go.This project has PostgreSQL database

## Create Request to the localhost:8080 And Show Response in the Terminal

1. add gorilla mux to the project

    go get github.com/gorilla/mux or go mod tidy

``*Note`` : go mod init

        if any error occurs : go: modules disabled by GO111MODULE=off; see 'go help modules'
        use this cmd in the terminal : export GO111MODULE=on

2. ``r.HandleFunc("/", getHome).Methods("GET")`` : This will send the Get request to the /.

3. ``ListenAndServ()`` : this method listen all request and serv for request

```go
package main

import (
 "fmt"
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
 fmt.Println("Home home")

}
```

Run : go run main.go

Browser : localhost:8080

Response : See terminal You will See "Home home"

## Let's See this Information in the html page

1.create ``assets/templates`` in the current directory

2.create home.html page inside this directory

```htm
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Home</title>
  </head>
  <body>
        <h1>Home home</h1>
  </body>
</html>

```

3.Add some lines inside main.go

```go
tmp, _ := template.New("home.html").ParseFiles("./assets/templates/home.html")

 err := tmp.Execute(w, nil)

 if err != nil {
  log.Println("Error Executing template : ", err)
  return
 }
```

4.Full Function for getting information about home

```go
func getHome(w http.ResponseWriter, r *http.Request) {
 tmp, _ := template.New("home.html").ParseFiles("./assets/templates/home.html")

 err := tmp.Execute(w, nil)

 if err != nil {
  log.Println("Error Executing template : ", err)
  return
 }

}
```

main.go

```go
package main

import (
 "fmt"
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

}
```

5.Request localhost:8080 you will get some Response

Run : go run main.go

Request : localhost:8080

Response : Home home

## File Structure

![p1](https://user-images.githubusercontent.com/37740006/122336614-f0dfb400-cf5e-11eb-8e1e-e5ed807b356d.png)

## Organize Code

1. We will  Remove our Route/handler/controller out of the main funcntion
2. For getting our Route we will create a method and call that method inside the main func.This will help to separate our code and it will gives us flexibility to the code.

3. main func code:

```go
func main() {

 r, err := handler.NewServer()
 if err != nil {
  log.Fatal("Handler not Found")
 }

 srv := &http.Server{
  Handler:      r,
  Addr:         "127.0.0.1:8080",
  WriteTimeout: 15 * time.Second,
  ReadTimeout:  15 * time.Second,
 }
 log.Fatal(srv.ListenAndServe())

}
```

4. ```handler.NewServer() :``` create a handler package and inside handler package create NewServer().This Method contains all the route/controller

5. handler/home.go : Create home.go package that contains all home related information or business logic if you want to add.

```go
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
```

6. handler/handler.go : create a handler package and inside handler package create NewServer().This Method contains all the route/controller

```go
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
```

5.Request localhost:8080 you will get some Response

Run : go run main.go

Request : localhost:8080

Response : Home home

main.go

```go
package main

import (
 "Go-Project-With-Postgres/handler"
 "log"
 "net/http"

 "time"
)

func main() {

 r, err := handler.NewServer()
 if err != nil {
  log.Fatal("Handler not Found")
 }

 srv := &http.Server{
  Handler:      r,
  Addr:         "127.0.0.1:8080",
  WriteTimeout: 15 * time.Second,
  ReadTimeout:  15 * time.Second,
 }
 log.Fatal(srv.ListenAndServe())

}
```

handler.go

```go
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
```

home.go

```go
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
```

home.html

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Home</title>
  </head>
  <body>
        <h1>Home home</h1>
  </body>
</html>
```

## File Structure So Far

![p2](https://user-images.githubusercontent.com/37740006/122342685-baa63280-cf66-11eb-937f-3365e9576f08.png)

## Add Desing to the Site Tailwindcss

1. First add static file loader in the handler.This will allow to use staic files js,css other design stuff

handler.go

```go
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
```

### [Follow This Link](https://themesberg.com/knowledge-center/tailwind-css/html)

## File Structure SO far

![p3Small](https://user-images.githubusercontent.com/37740006/122346540-e3c8c200-cf6a-11eb-84a2-b8d808217464.png)

## Connect with POSTGRES DB

1. Create a Method inside main.go file

```go
func newDBFromConfig() string {
 dbParams := " " + "user=postgres"
 dbParams += " " + "host=localhost"
 dbParams += " " + "port=5432"
 dbParams += " " + "dbname=practice"
 dbParams += " " + "password=0"
 dbParams += " " + "sslmode=disable"

 return dbParams
}
```

2. Create storage/postgres/postgres.go

```go
package postgres

import (
 "github.com/jmoiron/sqlx"
 _ "github.com/lib/pq"
 "github.com/pkg/errors"
)

type Storage struct {
 db *sqlx.DB
}

func NewStorage(dbstring string) (*Storage, error) {
 db, err := sqlx.Connect("postgres", dbstring)
 if err != nil {
  return nil, errors.Wrapf(err, "unable to connect to postgres '%s'", dbstring)
 }
 err = db.Ping()
 if err != nil {
  return nil, err
 }
 return &Storage{db: db}, nil
}

```

2. Call newDBFromConfig()

main.go

```go
package main

import (
 "Go-Project-With-Postgres/handler"
 "Go-Project-With-Postgres/storage/postgres"
 "log"
 "net/http"

 "time"
)

func main() {
 dbString := newDBFromConfig()
 store, err := postgres.NewStorage(dbString)

 if err != nil {
  log.Fatal(err)
 }

 r, err := handler.NewServer(store)
 if err != nil {
  log.Fatal("Handler not Found")
 }

 srv := &http.Server{
  Handler:      r,
  Addr:         "127.0.0.1:8080",
  WriteTimeout: 15 * time.Second,
  ReadTimeout:  15 * time.Second,
 }
 log.Fatal(srv.ListenAndServe())

}

func newDBFromConfig() string {
 dbParams := " " + "user=postgres"
 dbParams += " " + "host=localhost"
 dbParams += " " + "port=5432"
 dbParams += " " + "dbname=practice"
 dbParams += " " + "password=0"
 dbParams += " " + "sslmode=disable"

 return dbParams
}
```

3. Create storage.go file inside postgres folder

```go
package storage

// import "time"

type (
 User struct {
  ID        int32     `db:"id"`
  FirstName string    `db:"first_name"`
  LastName  string    `db:"last_name"`
  Username  string    `db:"username"`
  Email     string    `db:"email"`
  /* IsActive  bool      `db:"is_active"`
  IsAdmin   bool      `db:"is_admin"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"` */
 }
)
```

4. Update handler.go

```go
package handler

import (
 "Go-Project-With-Postgres/storage/postgres"
 "html/template"
 "net/http"

 "github.com/Masterminds/sprig"
 "github.com/gorilla/mux"
)

type (
 Server struct {
  templates *template.Template
  store     *postgres.Storage
 }
)

func NewServer(st *postgres.Storage) (*mux.Router, error) {

 s := &Server{
  store: st,
 }

 if err := s.parseTemplates(); err != nil {
  return nil, err
 }

 r := mux.NewRouter()
 r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))
 r.HandleFunc("/", s.getHome).Methods("GET")
 return r, nil
}

func (s *Server) parseTemplates() error {
 templates := template.New("templates").Funcs(template.FuncMap{
  "strrev": func(str string) string {
   n := len(str)
   runes := make([]rune, n)
   for _, rune := range str {
    n--
    runes[n] = rune
   }
   return string(runes[n:])
  },
 }).Funcs(sprig.FuncMap())

 tmpl, err := templates.ParseGlob("assets/templates/*.html")
 if err != nil {
  return err
 }
 s.templates = tmpl
 return nil
}
```
