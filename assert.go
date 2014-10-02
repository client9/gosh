package conduit

import (
	"errors"
	"os"
	"fmt"
)

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
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		os.Exit(1)
		return false, errors.New(msg)
	}

	return true, nil
}
