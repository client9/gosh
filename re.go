package gosh

import (
	"regexp"
)

// RegexpCompile compiles a new regexp pattern
// See regexp.Compile
func RegexpCompile(pattern string) (*regexp.Regexp, error) {
	return regexp.Compile(pattern)
}

// RegexpMatch returns true of target matches regexp
// See regexp.Match
func RegexpMatch(pattern string, target string) (bool, error) {
	return regexp.MatchString(pattern, target)
}

// RegexpQuoteMeta does regexp.QuoteMeta
func RegexpQuoteMeta(pattern string) string {
	return regexp.QuoteMeta(pattern)
}
