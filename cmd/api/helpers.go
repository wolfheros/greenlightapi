package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Retrieve the [id] URL parameter from current [Request Context}.
// all the parameters in request has been store in [Request Context]
// during the routing stage by [httprouter] frmaework

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// This method is use for sending json as response, it used parameters:
// [http.ResponseWriter], [HTTP status], [data] and [header map]

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// covert to json
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// append newline to json result
	js = append(js, '\n')

	// add header in the response
	for key, value := range headers {
		w.Header()[key] = value
	}

	// add the [Content-Type: application/json], and [status code], [json]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
