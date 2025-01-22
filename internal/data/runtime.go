package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Define an error that [UnmarshalJSON()] method can return
// If unable parse or convert the JSON string
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// Declare runtime type for custom json output, which mean it
// need to implement [MarshalJson()] interface [json.Marshaler] method
type Runtime int32

// implement [json.Marshaler] interface's [MarshalJSON()] method
func (r Runtime) MarshalJSON() ([]byte, error) {
	// formating the time with "mins" output the result
	jsonValue := fmt.Sprintf("%d mins", r)

	// add surrounding double-quotes to become a valid *JSON string*
	quotedJSONValue := strconv.Quote(jsonValue)

	// Convert the quoted string value to byte slice and return it
	return []byte(quotedJSONValue), nil
}

/**
* When Go decoding some JSON, it will check whether or not there is destination type
* satisfies the [json.Unmarshaler] interface, which mean implementation of [UnmarshalJSON()]
* method, use this implementation, we can determine how to decode the provide JSON into the
* target type, such as change the decoding behavior
*/

// Implement a [UnmarshalJSON()] method on the customer Runtime type so that it satifies the [json.Unmarshaler] interface
// we need modify the receiver, so we must use a pointer receiver. otherwise it will only modify the copy [Runtime] value
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error{
	// Incoming JSON value will be a string in the format "<runtime> mins"
	// First, removing the surrounding double-quotes from this string.
	// If cant unquote it, return [ErrInvalidRountimeFormat] error
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err!=nil {
		return ErrInvalidRuntimeFormat
	}

	//split the string based on " ", with string and number
	parts:= strings.Split(unquotedJSONValue, " ")

	// Check the output value which should two parts: "<number>" and "mins"
	if len(parts) > 2 {
		return ErrInvalidRuntimeFormat
	}

	// parse the first value to [int32]
	i, err:= strconv.ParseInt(parts[0], 10, 32)
	if err!=nil {
		return ErrInvalidRuntimeFormat
	}

	// here is using pointer,
	// deference the receiver with a new value
	*r = Runtime(i)

	return nil
}
