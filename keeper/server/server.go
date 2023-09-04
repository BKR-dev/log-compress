package server

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	mux "github.com/gorilla/mux"
)

var (
	port      = "9443"
	homeTempl = template.Must(template.New("").Parse(homeHTML))
)

func ServerStart() {
	fmt.Println("Starting Server: P: " + port)
	router := server()
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Create new muxServer from gorilla
func server() *mux.Router {
	r := mux.NewRouter()
	// use CORS middleware
	r.Use(mux.CORSMethodMiddleware(r))
	// define functions per route
	r.HandleFunc("/", homeHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/logs", logsHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions)
	r.HandleFunc("/times", timesHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions)
	s := r.PathPrefix("/archive/").Subrouter()
	s.HandleFunc("/", archivesHandler)
	s.HandleFunc("/{key}/", archiveHandler).Methods(http.MethodPost, http.MethodGet)

	return r
}

// /archive
func archivesHandler(w http.ResponseWriter, r *http.Request) {
	// returns overview of written archives

}

// /archive/{key}
func archiveHandler(w http.ResponseWriter, r *http.Request) {
	// akshually handle incoming file and write to db

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	var d = struct {
		Message string
		Time    time.Time
	}{
		"Hello",
		t,
	}
	err := homeTempl.Execute(w, d)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("logs directory: it werks!\n"))
}
func timesHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("times directory: it werks!\n"))
}

// use htmx for a nicer view

const homeHTML = `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>THE KEEPER</title>
    </head>
    <body>
				<h1>THE KEEPER</h1>
        {{.Message}}
				<br />
				{{.Time}}
    </body>
</html>`

const statusHTML = `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>THE KEEPER</title>
    </head>
    <body>
				<h1>THE KEEPER</h1>
        {{.Message}}
				<br />
				{{.Time}}
    </body>
</html>`
