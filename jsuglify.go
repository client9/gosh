package conduit

import (
	"bytes"
	"log"
	"os/exec"
)

// JSBeautify calls XYZ
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
