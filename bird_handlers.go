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
// convert the birds var into json 
birstListBytes, err := json.Marshal(birds)

//If there is an error, print it to the console, and return a server
// error response to the user
if err != nil {
	fmt.Println(fmt.Errorf("Error: %v", err))
	w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// write the JSON list of birds to the response
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
	birds = append(birds, bird)

	//Redirect user to HTML page using Redirect method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
