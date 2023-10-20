package main

import (
	"log"
	"net/http"
)

// handler - executing app logic
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))

}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displaying snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte("Create a snippet!"))
}

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
