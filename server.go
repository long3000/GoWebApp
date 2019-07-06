package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	// Port to URL
	Port = ":8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	response := "Current time: " + time.Now().String()
	fmt.Fprintln(w, response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/var/www/static.html")
}

func main() {
	http.HandleFunc("/static", serveStatic)
	http.HandleFunc("/", serveDynamic)
	// http.ListenAndServe(Port, http.FileServer(http.Dir("/var/www")))
	http.ListenAndServe(Port, nil)
}
