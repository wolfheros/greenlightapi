package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Create a [defered] function which will always running after panic
		defer func ()  {

			// Using builtin [recover()] function to check if there has a panic
			if err:= recover(); err!=nil {
				
				// If there was a panic,
				// Set a "Connection:close" header on the response
				// This will indicate the server close current connection after a response
				w.Header().Set("Connection", "close")


				// Builtin [recover()] function return a type any value
				// Using [fmt.Errorf()] funtion to normalise it into a error message.
				// Then using [helper.go] function [serverErrorResponse()] helper function to
				// log error, and send response to client [500 Internal Server Error]
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		// this line in here to make sense:
		// 1-> Need a return [http.Handler]
		// 2-> [http.Handler] is a interface the [ServeHTTP()] method has to be implemented.
		// 3-> This line won't be called until the request be servered in the [Server.ListenAndServe()] method
		next.ServeHTTP(w, r)
	})
}