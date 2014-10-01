package conduit

import (
	"bytes"
	"errors"
	"os/exec"
)

// JSHint calls out to jshint
// TODO accept command line args in pipeline
func JSHint(bytesin []byte) ([]byte, error) {
	cmd := exec.Command("/usr/bin/jshint", "-")
	cmd.Stdin = bytes.NewReader(bytesin)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out))
	}
	return bytesin, nil
}
