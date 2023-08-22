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

func server() *mux.Router {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/", homeHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/logs", logsHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions)
	r.HandleFunc("/times", timesHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions)
	return r
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
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logs directory: it werks!\n"))
}
func timesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("times directory: it werks!\n"))
}

const homeHTML = `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>THE KEEPER</title>
    </head>
    <body>
        {{.Message}}
				{{.Time}}
    </body>
</html>
`
