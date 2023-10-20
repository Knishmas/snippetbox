package main

import (
	"log"
	"net/http"
)

func main() {
	//init a new servemux
	mux := http.NewServeMux()
	//router - maps url with corresponding handler
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	//creating a new server.
	//Parameters: (TCP network address, your servemux)
	err := http.ListenAndServe(":4000", mux)
	//If there's an error, then it's logged.
	log.Fatal(err)
}
