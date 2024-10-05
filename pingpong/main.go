package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var pongNumber = 0

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong %v", pongNumber)
		pongNumber++
	}).Methods("GET")

	http.ListenAndServe(":80", r)
}
