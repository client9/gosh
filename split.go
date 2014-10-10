package conduit

import (
	"errors"
	"strings"
)

// Split splites text on newline
//
// split [spliton string] arg []byte|string
func Split(args ...interface{}) ([]string, error) {
	spliton := "\n"
	idx := 0
	if len(args) == 0 || len(args) > 2 {
		return nil, errors.New("Wrong number of args")
	}
	if len(args) == 2 {
		spliton = String(args[0])
		idx = 1
	}
	source := strings.Trim(String(args[idx]), spliton)
	if len(source) == 0 {
		return []string{}, nil
	}
	lines := strings.Split(source, spliton)
	return lines, nil
}

// Join merges an []string into a single string
// arg0 is joinwith string
func Join(args ...interface{}) (string, error) {
	joinwith := "\n"
	idx := 0
	if len(args) == 0 || len(args) > 2 {
		return "", errors.New("Wrong number of args")
	}
	if len(args) == 2 {
		joinwith = String(args[0])
		idx = 1
	}
	ary, err := StringArray(args[idx])
	if err != nil {
		return "", err
	}
	return strings.Join(ary, joinwith), nil

}
