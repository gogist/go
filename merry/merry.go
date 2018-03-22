package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ansel1/merry"
	"github.com/hashicorp/hcl/hcl/parser"
)

var InvalidInputs = errors.New("Input is invalid")

func main() {
	// create a new error, with a stacktrace attached
	err := merry.New("bad stuff happened")

	// create a new error with format string, like fmt.Errorf
	err = merry.Errorf("bad input: %v", os.Args)

	// capture a fresh stacktrace from this callsite
	err = merry.Here(InvalidInputs)

	// Make err merry if it wasn't already.  The stacktrace will be captured here if the
	// error didn't already have one.  Also useful to cast to *Error
	err = merry.Wrap(err)

	// override the original error's message
	err.WithMessagef("Input is invalid: %v", os.Args)

	// Use Is to compare errors against values, which is a common golang idiom
	merry.Is(err, InvalidInputs) // will be true

	// associated an http code
	err.WithHTTPCode(400)

	_, perr := parser.Parse([]byte("blah"))
	err = merry.Wrap(perr)
	// Get the original error back
	merry.Unwrap(err) // will be true
	fmt.Println(err)

	// Print the error to a string, with the stacktrace, if it has one
	s := merry.Details(err)
	fmt.Println("detail", s)

	// Just print the stacktrace (empty string if err is not a RichError)
	// s := merry.Stacktrace(err)

	// Get the location of the error (the first line in the stacktrace)
	file, line := merry.Location(err)
	fmt.Println("location", file, line)

	// Get an HTTP status code for an error.  Defaults to 500 for non-nil errors, and 200 if err is nil.
	code := merry.HTTPCode(err)
	fmt.Println("code", code)

}
