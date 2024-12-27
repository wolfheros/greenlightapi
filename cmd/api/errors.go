package main

import (
	"fmt"
	"net/http"
)

// log any error happend on the server
func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)


}

// Response with JSON result
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any){
	env := envelope{"error":message}

	// Write response use [helper] package writeJSON() helper function
	err:= app.writeJSON(w, status, env, nil)
	if err!=nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// Response with server error
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error){
	app.logError(r, err)

	message:= "the server encountered a problem and could not proces your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// Response 404 Not Found 
func(app *application) notFoundResponse(w http.ResponseWriter, r *http.Request){
	message:="the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// Response 405 methodNotAllowedResponse()
func(app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request){
	message:= fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}


