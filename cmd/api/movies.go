package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.wolfheros.com/internal/data"
	"greenlight.wolfheros.com/internal/validator"
)

// create movie handler
// response to [POST /v1/movies] endpoint
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	
	// Declare a anonymous struct hold the information from client request body
	// The struct will be the [target decode destination]
	// 
	var input struct{
		Title string `json:"title"` // struct tags
		Year int32 `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres []string `json:"genres"`	// slice
	}

	// Old version:
	// Initialize a new [json.Decoder] instance which reads from the request body
	// Using [Decode()] method to decode the body contents into the struct.
	// Send back [400 Bad Request] if there are error happened during decoding
	// Passing [input] pointer to the [Decode()] method as target decode destination
	// err:= json.NewDecoder(r.Body).Decode(&input)
	
	// New version:
	// Decode the result to anonymous struct
	err:=app.readJSON(w, r, &input)
	if err!=nil {
		// app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		app.badRequestResponse(w, r, err)
		return
	}

	movie:= &data.Movie{
		Title: input.Title,
		Year: input.Year,
		Runtime: input.Runtime,
		Genres: input.Genres,
	}

	// Initialize a new Validator instance to verify the client import
	v:=validator.New()

	// At the end check is there any failed validation by checking validator instance
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w,r,v.Errors)
		return
	}

	// Dumping the contents of the input struct in a http response
	// fmt.Fprintln(w, "create a new movie")
	fmt.Fprintf(w, "%+v\n", input)
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
