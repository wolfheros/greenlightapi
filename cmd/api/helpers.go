package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Define an envelope type for enveloping the data result.
type envelope map[string]any

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

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// covert to json, use [json.MarshalIndent()]
	// function add whitespace to the output, use for each element
	// noline prefix [""]
	// table indents ["\t"]
	js, err := json.MarshalIndent(data, "", "\t")
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

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error{

	// Limits the max size of request body to 1MB by use [http.MaxBytesReader()]
	maxBytes:= 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Decode the request body
	// err:= json.NewDecoder(r.Body).Decode(dst)

	// Create a new Decoder, then setting decoder disallow any unknow field while mapping it to strcut.
	dec:= json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err:= dec.Decode(dst)

	if err!=nil {
		
		// if there is an error during decode
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		// add a new [maxBytesError] variable
		var maxBytesError *http.MaxBytesError

		switch{
		
		// Use [errors.As()] function to check whether the error has the type *json.SyntaxError
		// return error message and the location of problem
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contain badly-formed JSON (at charact %d)", syntaxError.Offset)

		// check is or not a [io.ErrUnexpectedEOF]
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contain badly-formed JSON")

		// catch any *json.UnmarshalTypeError errors
		case errors.As(err,&unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for filed %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		
			// if the body is Empty
		// it will return a EOF error
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		
		// If the JSON contains a field which cnannot to be mapped to the target destination
		// then [Decoder()] will now return an error message in the format ["json: unknown field"]
		// there is a disccution about go try to take this error to a distinct error type in the future.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fildName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fildName)

		// Use the [errors.As()] check whether the error has the type [http.MaxBytesError]
		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)
		// [json.InvalidUnmarshalError] error will be returned
		// when a invalid arguments pass to [Decode()]
		// panic VS return
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		// Default return any other error
		default:
			return err
		}
	}

	// call [Decode()] again to make sure there only a single JSON value,
	// Because the Decode() can unload multiple values in a single request.Body
	// if it return [EOF] which mean it is the end of the body, otherwise there is more 
	// data, which in this situation we won't need it, and see it as Error.
	// Here:
	// ---> Using a empty anonymous struct's pointer as the destination
	
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must only contain a single JSON value per request")
	}
	return nil
}
