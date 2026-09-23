package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	/*
		log.new() to create custom logger for writting information messages
		Parameters:
				- destination to write the logs to (os.Stdout)
				- string prefix message (INFO followed by a tab)
				- flags to indicate what additional information to include joined with bitwise OR operator |
	*/
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	/*
		log.new() to create custom logger for error messages, but with Stderr
		Parameters:
				- destinations to write the logs to (os.Stderr)
				- string prefix message (ERROR followed by a tab)
				- flags to include relevant file name and line number use log.Lshofrtfile flag
	*/
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	/*
		Initialize a new http.Server struct.
		set the Addr and Handler fields so that the server uses the same
		network address and routes as before.
		set ErrorLog field: server now uses the custom errorLog logger

	*/
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
