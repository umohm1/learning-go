package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird 

func getBirdHandler(w http.ResponseWriter, r *http.Request) {

	/*
		The list of birds is now taken from the store instead of the package level  `birds` variable we had earlier

		The `store` variable is the package level variable that we defined in
		`store.go`, and is initialized during the initialization phase of the
		application
	*/

  birds, err := store.GetBirds()

	// Everything else is the same as before
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	//Create a new instance of bird
	bird := Bird{}

	// All data is sent as HTML form data. ParseForm parses the form values
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf(Error: %v, err))
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	//Get bird information from form info
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	//Append existing list of birds w/ a new entry
	err = store.CreateBird(&bird)
	if err != nil {
		fmt.Println(err)
	}

	//Redirect user to HTML page using Redirect method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
