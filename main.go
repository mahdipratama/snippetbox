package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// 'Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	/*
		Important: return from handler, or it would keep executing
		and also wriite the "Hello from Snippetbox" message
	*/

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snipppet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		/*
			Use the Header().Set() method to add an 'Allow: POST' header
			to the response header map. The first parameter is the header name,
			and the second parameter is the header value.
		*/
		w.Header().Set("Allow", http.MethodPost)

		/*
			Use the http.Error() function to send a 405 status code
			and "Method Not Allowed" string as the response body.
		*/
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Display a new snipppet..."))
}

func main() {
	/*
		Use the http.NewServeMux() function to initialize a new servemux, then
		register the home function as the handler for the "/" URL pattern.
		Go servemux treats the URL pattern "/" like a catch-all.
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	/*
		User the http.ListenAndServer() function to start a new web server.
		Pass in two parameters:
			- The TCP network address to listen on :4000
			- The servemux we just created

		If http.ListenAndServe() returns an error,
		we user the log.Fatal() function to log the error message and exit.
		Note: that any error returned by http.ListenAndServer() is always non-nil
	*/
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
