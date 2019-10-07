package main

import "database/sql"

// Need the sql library to interact w/ the DB

// The store has 2 methods. One to add a new bird and another to get all existing birds. Each method returns an error in case of failure
type Store interface {
	CreateBird(bird *Bird) error
	GetBirds() ([]*Bird, error)
}

// The dbStore struct implements the Store interface. It takes the sql DB connection object, which represents the database connection
type dbStore struct {
	db *sql.DB
}

// 'Bird' is a simple struct which has "species" and "description" attributes
// The first underscore means that we don't care about what's returned from
// this insert query. We just want to know if it was inserted correctly,
// and the error will be populated if it wasn't
func (store *dbStore) CreateBird(bird *Bird) error {
	_, err := store.db.Query("INSERT INTO birds(species, description) VALUES ($1, $2)", bird.Species, bird.Description)
	return err
}

func (store *dbStore) GetBirds() ([]*Bird, error) {
	// Query the database for all birds, and return the result to the
	// `rows` object
	rows, err := store.db.Query("SELECT species, description from birds")
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of birds
	birds := []*Bird{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		bird := &Bird{}
		// Populate the `Species` and `Description` attributes of the bird, and return incase of an error
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for the next row
		birds = append(birds, bird)
	}
	return birds, nil
}

// The store variable is a package level variable that will be available for
// use throughout our application code
var store Store

/*
We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock
*/
func InitStore(s Store) {
	store = s
}
