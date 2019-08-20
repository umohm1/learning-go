package main

import (
	// formats operations
	"fmt"
	// implements clients and servers
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Declare a new router
	r := mux.NewRouter()

	r.HandleFunc("/hello", handler).Methods("GET")

	// listen and serve on port 8080
	http.ListenAndServe(":8080", r)

}

// must follow func signature of ResponseWriter and Request type
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
