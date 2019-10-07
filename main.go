package main

import (
	// formats operations
	"fmt"
	// implements clients and servers
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// create a new router and return it outside of main func
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")

	// The fileserver is wrapped in the `stripPrefix` method to
	// remove the "/assets/" prefix when looking for files.
	// w/o it the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileDirectory := http.Dir("./assets")
	staticFileHandler := http.StripPrefix("/assets", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
	fmt.Println("Starting server...")
	connString := "dbname=temp sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	// listen and serve on port 8080
	r := newRouter()
	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", r)

}

// must follow func signature of ResponseWriter and Request type
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
