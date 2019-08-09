package main

import (
	// formats operations
	"fmt"
	// implements clients and servers
	"net/http"
)

func main() {
	// method accepts path and function
	http.HandleFunc("/", handler)

	// listen and serve on port 8080
	http.ListenAndServe(":8080", nil)

}

// must follow func signature of ResponseWriter and Request type
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
