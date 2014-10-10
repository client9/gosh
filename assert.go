package conduit

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// DevNull maps any input to empty string
// Useful in validators
func DevNull(data interface{}) string {
	return ""
}

// Array converts a series or arguments to an []string
//
func Array(data ...string) []string {
	return data
}

// Integer attempt to convert an input into a golang 'int' type
func Integer(data interface{}) (int, error) {
	var val int
	var err error

	switch data.(type) {
	case int:
		val = data.(int)
	case uint:
		val = int(data.(uint))
	case int32:
		val = int(data.(int32))
	case int64:
		val = int(data.(int64))
	case string:
		val, err = strconv.Atoi(data.(string))
		if err != nil {
			return 0, fmt.Errorf("Unable to convert %q to integer", data)
		}
	default:
		return 0, fmt.Errorf("Unable to convert %T to integer", data)
	}
	return val, nil
}

// StringArray converts input into a string array... this is kinda bogus
//
func StringArray(data interface{}) ([]string, error) {
	switch data.(type) {
	case []byte:
		raw, _ := data.([]byte)
		return strings.Split(string(raw), "\n"), nil
	case []string:
		source := data.([]string)
		return source, nil
	default:
		return nil, errors.New("unable to convert to []string")
	}

}

// String attempts to convert input into a string representation
func String(data interface{}) string {
	switch data.(type) {
	case string:
		bits, _ := data.(string)
		return bits
	case []byte:
		bits, _ := data.([]byte)
		return string(bits)
	default:
		return fmt.Sprintf("%v", data)
	}
}
