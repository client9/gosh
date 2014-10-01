package conduit

import (
	"bytes"
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
