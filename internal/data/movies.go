package data

import (
	"time"
	"greenlight.wolfheros.com/internal/validator"
)

/*
*
The [Movies] struct is store data like this:

	{
		"id": 123,
		"title": "Casablanca"
		"runtime": 102,
		"genres": [
			"drama",
			"romance",
			"war",
		],
		"version": 1
	}

**
*/
type Movie struct {
	ID        int64     `json:"id"` 				// struct tags
	CreatedAt time.Time `json:"-"` 					// use [-] (hyphen) to remove from the result
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"` 	// Add the omitempty directive to remove the value while its [empty], [nil] or [""]
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"` 	//use slice to store multiple value.
	Version   int32     `json:"version"`          	//version start as 1 will increase based on the update operation
}



// Using [Check()] method to execute each different verification
// [Check()] first param is [bool] value, it will decide the behaviour of 
// [Check()] method
func ValidateMovie(v *validator.Validator, movie *Movie){
	// check title
	v.Check(movie.Title!="", "title", "must be provided")
	// check title length less than 500
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
	// check year valid
	v.Check((movie.Year != 0) && (movie.Year >=1888) && (movie.Year <= int32(time.Now().Year())), "year", "must be valid value, over 1888 but not in the future")
	// check runtime value
	v.Check(movie.Runtime > 0, "runtime", "must be provided and must be a positive integer")
	// check genres value
	v.Check(movie.Genres != nil && len(movie.Genres) >=1 && len(movie.Genres) <= 5, "genres" , "Must be provided and at least 1 genres less than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "Must not contain duplicate values")
}
