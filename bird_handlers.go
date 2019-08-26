type Bird struct {
	Species string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird 

func getBirdHandler(w http.ResponseWriter, r *http.Request)
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
