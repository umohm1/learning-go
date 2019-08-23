package main

import (
	// formats operations
	"fmt"
	// implements clients and servers
	"net/http"
	"github.com/gorilla/mux"
)

// create a new router and return it outside of main func
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func main() {
	// listen and serve on port 8080
	r := newRouter()
	http.ListenAndServe(":8080", r)

}

// must follow func signature of ResponseWriter and Request type
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}