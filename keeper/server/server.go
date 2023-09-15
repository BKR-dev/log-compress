package server

import (
	"fmt"
	"local/db"
	"local/model"
	"log"
	"net/http"
	"text/template"
	"time"

	mux "github.com/gorilla/mux"
)

var (
	port        = "9443"
	homeTempl   = template.Must(template.New("").Parse(homeHTML))
	statusTempl = template.Must(template.New("").Parse(statusHTML))
	dB, _       = db.GetDB()
	qS          = db.NewQueryService(dB)
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
	r.HandleFunc("/", homeHandler).
		Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/logs", logsHandler).
		Methods(http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions)
	r.HandleFunc("/times", timesHandler).
		Methods(http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions)
	subArch := r.PathPrefix("/archive/").Subrouter()
	subArch.HandleFunc("/", archivesHandler)
	subArch.HandleFunc("/{key}/", archiveHandler).
		Methods(http.MethodPost, http.MethodGet)
	r.HandleFunc("/status", statusHandler).
		Methods(http.MethodGet)

	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	var d = struct {
		Message string
		Time    string
	}{
		"Hello",
		time.DateTime,
	}
	err := homeTempl.Execute(w, d)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}
}

// /archive
func archivesHandler(w http.ResponseWriter, r *http.Request) {
	// returns overview of written archives

}

// /archive/{key} key == id? hostname? applicationname?
func archiveHandler(w http.ResponseWriter, r *http.Request) {
	// akshually handle incoming file and write to db

}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logs directory: it werks!\n"))
}

func timesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("times directory: it werks!\n"))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	entries := trfmLstr(qS.GetAllLogEntriesWithGorm())
	fmt.Println(entries[0])
	var d = struct {
		Entries []model.LogString
	}{
		entries,
	}
	err := statusTempl.Execute(w, d)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
	}

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
				<ul>
        {{range .Entries}}
					<li class="hostname">{{.Hostname}}</li>
					<li class="appname">{{.ApplicationName}}</li>
				{{end}}
				<br/>
				</ul>
    </body>
</html>`
