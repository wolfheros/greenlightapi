package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// initialize a new router instance, 
	// [router] has implemented [http.Handler] interface
	router := httprouter.New()

	//Convert the [notFoundResponse()] helper to a [http.Handler]
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	//Convert the methodAllowedResponse() helper to a [http.Handler]
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Register route to the router:
	// -> Request Methods
	// -> URL patterns
	// -> Handler functions
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)


	// return router
	return app.recoverPanic(router)
}
