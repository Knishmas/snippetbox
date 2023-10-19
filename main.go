package main

import (
	"log"
	"net/http"
)

// handler - executing app logic
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	//init a new servemux
	mux := http.NewServeMux()
	//router - maps url with corresponding handler
	mux.HandleFunc("/", home)

	log.Print("Starting server on :4000")
	//creating a new server.
	//Parameters: (TCP network address, your servemux)
	err := http.ListenAndServe(":4000", mux)
	//If there's an error, then it's logged.
	log.Fatal(err)
}
