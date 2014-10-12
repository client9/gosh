package gosh

import (
	"bytes"
	"errors"
	"os/exec"
)

// StandardCLI handles simulating a normal command line environment
// In particular, if the last arg is a []string it is merged
// into command line
func StandardCLI(exe string, extra []string, args []interface{}) ([]byte, error) {
	flags := make([]string, 0, len(args)+len(exe))

	var bytesin []byte

	if extra != nil {
		for _, val := range extra {
			flags = append(flags, val)
		}
	}
	lastarg := args[len(args)-1]
	switch lastarg.(type) {
	case []byte:
		bytesin, _ = args[len(args)-1].([]byte)
		// get flags.. convert them all to strings
		for _, val := range args[:len(args)-1] {
			flags = append(flags, String(val))
		}
	case []string:
		filenames, _ := lastarg.([]string)
		for _, val := range args[:len(args)-1] {
			flags = append(flags, String(val))
		}
		for _, val := range filenames {
			flags = append(flags, val)
		}
	default:
		for _, val := range args {
			flags = append(flags, String(val))
		}
	}
	cmd := exec.Command(exe, flags...)
	if bytesin != nil {
		cmd.Stdin = bytes.NewReader(bytesin)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out))
	}
	return out, nil
}
