package conduit

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

// JSBeautify calls XYZ
// TODO add command line args to pipeline
func JSBeautify(bytesin []byte) []byte {
	cmd := exec.Command("/usr/bin/js-beautify", "-f", "-")
	cmd.Stdin = bytes.NewReader(bytesin)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("jsbeauty failed: %s", err)
	}
	return out
}

// JSUglify calls XYZ
// TODO add command line args to pipeline
func JSUglify(bytesin []byte) []byte {
	cmd := exec.Command("/usr/bin/uglifyjs", "-")
	cmd.Stdin = bytes.NewReader(bytesin)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("uglifyjs failed: %s", err)
	}
	return out
}

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
