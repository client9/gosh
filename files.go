package conduit

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Glob does filesystem matching
// returns files []string
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

// Open is similar to ioutil.ReadFile
//
// TODO: is this really "cat"? or is Open better?
//  Or ReadFile ala ioutil.ReadFile
//
func Open(fname string) ([]byte, error) {
	fbytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return fbytes, nil
}

// Write similar to ioutil.WriteFile
// TODO PERMISSIONS
func Write(fname string, bytesin []byte) (string, error) {
	err := ioutil.WriteFile(fname, bytesin, 0777)
	if err != nil {
		return "", err
	}
	return fname, nil
}

/*
type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)
}
*/

// FileStat return os.FileInfo object or error
func FileStat(fname string) (os.FileInfo, error) {
	return os.Stat(fname)
}

// TimeNow return the current time.Time
func TimeNow() time.Time {
	return time.Now()
}

// FileModTime return the time of when the file was modified
// An alias for  (stat "filename").ModTime
//
func FileModTime(fname string) (time.Time, error) {
	fi, err := os.Stat(fname)
	if err != nil {
		return time.Time{}, err
	}
	return fi.ModTime(), nil
}
