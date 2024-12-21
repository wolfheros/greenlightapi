package data

import "time"

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
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // use [-] (hyphen) to remove from the result
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"` // Add the omitempty directive to remove the value while its [empty], [nil] or [""]
	Runtime   Runtime     `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"` //use slice to store multiple value.
	Version   int32     `json:"version"`          //version start as 1 will increase based on the update operation
}
