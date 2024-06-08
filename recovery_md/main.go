package main

import (
	"log"
	"net/http"
)

func regular(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Regular"))
}

func panics(w http.ResponseWriter, r *http.Request) {
	panic("Something went wrong!")
}

func latePanic(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Some initial text"))
	panic("Something went terribly wrong!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/regular", regular)
	mux.HandleFunc("/panics", panics)
	mux.HandleFunc("/latepanic", latePanic)
	log.Fatal(http.ListenAndServe(":3000", recoveryMw(mux, true)))
}
