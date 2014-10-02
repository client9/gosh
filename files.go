package conduit

import (
	"path/filepath"
	"io/ioutil"
)

func Glob(pat string) ([]string, error) {
	matches, err := filepath.Glob(pat)
	if err != nil {
		return nil, err
	}
	if matches == nil {
		return []string{}, nil
	}
	return matches, nil
}

func Open(fname string) ([]byte, error) {
	fbytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return fbytes, nil
}
	
