package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	mux "github.com/gorilla/mux"
)

func main() {

}

var (
	// filebeat output config
	// path: filePath
	// filename: filebeat
	// rotate_every_kb: fileSize
	// permissions: 0600
	// rotate_on_startup: false
	filePath     = "/tmp/filebeat"
	fileBasename = "filebeat"
	fileSize     = 100 * 1024 * 1025 // 100mb
	keeperURL    = "not_defined_yet"
)

func readFile(filename string) (*[]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	file.Close()
	buf := make([]byte, fileSize)
	_, err = io.ReadFull(file, buf)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

// Create new muxServer from gorilla
func senderServer() *mux.Router {
	r := mux.NewRouter()
	// use CORS middleware
	r.Use(mux.CORSMethodMiddleware(r))
	// define functions per route
	r.HandleFunc("/status", statusHandler).Methods(http.MethodGet)
	return r
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// add htmx view for status of the sender
}

// URL where to send and byte slice of data
func sendFileToKeeper(url string, body []byte) {
	r, err := http.NewRequest(http.MethodPost, keeperURL, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}

	r.Header.Add("Content-Type", "Data")

}
