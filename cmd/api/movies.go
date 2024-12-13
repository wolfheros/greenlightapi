package main

import (
	"fmt"
	"net/http"
)

// create movie handler
// response to [POST /v1/movies] endpoint
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// show movie handler
// response to [GET /v1/movie/:id] endpoint
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// get [id] param from it
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
