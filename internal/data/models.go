package data

import (
	"database/sql"
	"errors"
)

//Define a custom Error for [Get()] method
var(
	ErrRecordNotFound = errors.New("record not found")
)

// Create a Models struct which will wrap all the Models in the future, include MovieModels
type Models struct{
	Movies MovieModel
}

// Create [Models] instance
func NewModels(db *sql.DB)Models{
	return Models{
		Movies: MovieModel{DB: db},
	}
}