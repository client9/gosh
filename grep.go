package conduit

import (
	"bytes"
	"errors"
	"path"
	"regexp"
	"strings"
)

// Grep pattern (flags) []string
//
func Grep(args ...interface{}) ([]string, error) {
	idx := 0
	reflags := ""
	invert := false

	switch len(args) {
	case 2:
		idx = 0
	case 3:
		idx = 1
		flagparts := strings.Split(String(args[0]), ",")
		for _, val := range flagparts {
			switch val {
			case "invert":
				invert = true
			case "insensitive":
				reflags = reflags + "i"
			case "multiline":
				reflags = reflags + "m"
			case "dotall":
				reflags = reflags + "s"
			case "ungreedy":
				reflags = reflags + "U"
			}
		}
	default:
		return nil, errors.New("wrong args")
	}

	pat := String(args[idx])
	if len(reflags) != 0 {
		pat = "(?" + reflags + ")" + pat
	}
	patre, err := regexp.Compile(pat)
	if err != nil {
		return nil, err
	}
	lines, err := StringArray(args[idx+1])
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		ok := patre.MatchString(line)
		if (ok && !invert) || (!ok && invert) {
			out = append(out, line)
		}
	}
	return out, err
}

// PathMatch does path-style glob matching per line
// this of this as a glob-style grep
func PathMatch(args ...interface{}) ([]string, error) {
	if len(args) != 2 {
		return nil, errors.New("wrong args")
	}
	pat := String(args[0])
	lines, err := StringArray(args[1])
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(args))
	for _, val := range lines {
		ok, err := path.Match(pat, val)
		if err != nil {
			return nil, err
		}
		if ok {
			out = append(out, val)
		}
	}
	return out, nil
}

// LineCount counts number of lines in an array
func LineCount(bytesin []byte) int {
	count := 0
	for len(bytesin) > 0 {
		idx := bytes.IndexByte(bytesin, '\n')
		if idx == -1 {
			return count
		}
		count++
		bytesin = bytesin[idx+1:]
	}
	return count
}
