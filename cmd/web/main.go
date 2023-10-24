package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()
	//Custom loggers | info & error for leveled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//init a new servemux
	mux := http.NewServeMux()
	//router - maps url with corresponding handler
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Print("Starting server on %s", *addr)
	//creating a new server.
	//Parameters: (TCP network address, your servemux)
	err := http.ListenAndServe(*addr, mux)
	//If there's an error, then it's logged.
	errorLog.Fatal(err)
}
