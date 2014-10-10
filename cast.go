package conduit

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// MapAny is a string-to-any mapping
type MapAny map[string]interface{}

// Set sets an entry in the map
func (m MapAny) Set(key string, value interface{}) MapAny {
	m[key] = value
	return m
}

// Get returns an entry in the map
func (m MapAny) Get(key string) interface{} {
	return m[key]
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
