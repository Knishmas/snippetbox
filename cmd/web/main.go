package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// custom struct for application wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()
	//Custom loggers | info & error for leveled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	//init a new servemux
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//router - maps url with corresponding handler
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	//new http.Server struct so that http.server error uses our custom logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Print("Starting server on %s", *addr)
	//creating a new server.
	//Parameters: (TCP network address, your servemux)
	err := srv.ListenAndServe()
	//If there's an error, then it's logged.
	errorLog.Fatal(err)
}
