package data

import (
	"fmt"
	"strconv"
)

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
