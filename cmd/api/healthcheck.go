package main

import (
	"net/http"
)

// handler register as application's method
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	// Careful the string literal (enclosed with backticks)
	// so that include double-quotes characters in the JSON without needing to escape it

	// Creating a [map[string]string] which hold the [JSON] string
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "the server encoutered a problem and could not process your request", http.StatusInternalServerError)
	}
}
