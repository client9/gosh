package conduit

import (
	"errors"
	"fmt"
	"os"
)

// Type returns the native go type as a string.
//
// Arguments: 1
// Input: anything
// Output: string
// Error: never
//
func Type(arg interface{}) string {
	return fmt.Sprintf("%T", arg)
}

// Debug writes a message to stderr and returns the original input
func Debug(msg string, arg interface{}) interface{} {
	os.Stderr.WriteString(msg)
	return arg
}

// Assert halts execution if argument is false
// If two arguments, then first argument is message
func Assert(args ...interface{}) (bool, error) {
	var truth bool
	var ok bool
	var msg string
	var idx int

	switch len(args) {
	case 1:
		idx = 0
		msg = "Assertion failed"
	case 2:
		idx = 1
		msg, ok = args[0].(string)
		if !ok {
			return false, errors.New("Expected message for arg 0")
		}
	default:
		return false, errors.New("Got too many arguments")
	}
	truth, ok = args[idx].(bool)
	if !ok {
		return false, errors.New("Expected boolean value")
	}

	if !truth {
		//fmt.Fprintf(os.Stderr, "%s", msg)
		return false, errors.New(msg)
	}

	return true, nil
}
