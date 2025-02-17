package validator

import (
	"regexp"
	"slices"
)

// Declare a regular expression for snity checking the format of email address
var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)


// Define a new Validator type which contain a map of validation errors
type Validator struct{
	Errors map[string]string
}

// [New()] method is a helper create a instance of [Validator] struct
// Be aware here using [make()] to create a empty [map]
// Here the return is a [Validator] pointer
func New() *Validator{
	return &Validator{Errors: make(map[string]string)}
}

// [Validator] struct methods, return ture if it doesn't contain any entries
func (v *Validator) Valid() bool{
	return len(v.Errors) == 0
}

// Add errors to the map
func (v *Validator) AddError(key, message string) {
	if _, exists:= v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message to the map only if a validation check is not [ok]
func (v *Validator) Check(ok bool, key, message string){
	if !ok {
		v.AddError(key, message)
	}
}

// Generic function which returns true if a specific value in a list of permitted value.
func PermittedValue[T comparable](value T, permittedValues ...T) bool{
	return slices.Contains(permittedValues, value)
}

// Matches returns true if a string value matches a specific regxp pattern
func Matches(value string, rx *regexp.Regexp) bool{
	return rx.MatchString(value)
}


// Generic function which returns true if all values in a slice are unique
func Unique[T comparable](values []T) bool{
	uniqueValues := make(map[T]bool)

	// because slice can storage duplicate value, but not [map]
	// use this mechanic to remove duplicate value.
	for _, value:= range values{
		uniqueValues[value] = true
	}
	
	return len(values) == len(uniqueValues)
}