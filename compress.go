package gosh

import (
	"bytes"
	"compress/gzip"
)

// Gzip compress input:
func Gzip(bytesin []byte) ([]byte, error) {
	out := bytes.Buffer{}
	zip := gzip.NewWriter(&out)
	_, err := zip.Write(bytesin)
	if err != nil {
		return nil, err
	}
	err = zip.Close()
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// Gunzip decompresses input via Gzip
func Gunzip(bytesin []byte) ([]byte, error) {
	out := bytes.Buffer{}
	zip, err := gzip.NewReader(&out)
	if err != nil {
		return nil, err
	}
	_, err = zip.Read(bytesin)
	if err != nil {
		return nil, err
	}
	err = zip.Close()
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
