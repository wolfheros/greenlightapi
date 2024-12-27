package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.wolfheros.com/internal/data"
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
		// http.NotFound(w, r)
		app.notFoundResponse(w, r)
		return
	}

	// create a [Movies] struct instance, contain some dummy data
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablance",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		// app.logger.Error(err.Error())
		// http.Error(w, "The server encouted a problem and could not process your request", http.StatusInternalServerError)
		app.serverErrorResponse(w, r, err)
	}

	// fmt.Fprintf(w, "show the details of movie %d\n", id)
}
