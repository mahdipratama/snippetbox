package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// command-line flag named 'addr'
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Parse the command-line flag then assigns it to the addr variable
	// call this before use the 'addr' or it'll still contain the default value ':4000'
	// if any error occur during parsing, the app will be terminated
	flag.Parse()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// value of 'addr' returned from the flag.String()
	// is a pointer to the flag value, not the value itself
	// need to dereference ( * symbol ) the pointer before use it.
	// USAGE: go run ./cmd/web -addr=":<PORT>"
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)

}
