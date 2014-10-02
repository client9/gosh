package conduit

import (
	"os/exec"
	"bytes"
	"errors"
)

/*
func Goimports(bytesin []byte) ([]byte, error) {
        cmd := exec.Command("goimports")
        cmd.Stdin = bytes.NewReader(bytesin)
        out, err := cmd.CombinedOutput()
        if err != nil {
                return nil, errors.New(string(out))
        }
        return out, nil
}
*/

func Goimports(args ...interface{}) ([]byte, error) {

	flags := make([]string, len(args)-1)
	for pos, val := range args[:len(args)-1] {
		flags[pos] = String(val)
	}

	bytesin, ok := args[len(args)-1].([]byte)
	if !ok {
		return nil, errors.New("Wrong type, expected []bytes")
	}
        cmd := exec.Command("goimports", flags...)
        cmd.Stdin = bytes.NewReader(bytesin)
 	out, err := cmd.CombinedOutput()
        if err != nil {
                return nil, errors.New(string(out))
        }
        return out, nil	
}

