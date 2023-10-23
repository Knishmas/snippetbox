package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	//init a new servemux
	mux := http.NewServeMux()
	//router - maps url with corresponding handler
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on %s", *addr)
	//creating a new server.
	//Parameters: (TCP network address, your servemux)
	err := http.ListenAndServe(*addr, mux)
	//If there's an error, then it's logged.
	log.Fatal(err)
}
