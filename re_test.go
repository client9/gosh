package gosh

import (
	"fmt"
	"testing"
)

func TestReStrings(t *testing.T) {
	var casetests = []struct {
		script   string
		input    string
		expected string
	}{
		{` $ | regexpquotemeta`, "[]", "\\[\\]"},
		{` $ | regexpmatch "foo"`, "xxxfooyyyy", "true"},
		{` $ | (regexpcompile "foo").MatchString`, "xxxxfooozzzz", "true"},
	}

	for pos, tt := range casetests {

		out, err := Execute(fmt.Sprintf("testcase%d", pos), tt.script, tt.input)
		if err != nil {
			panic(err)
		}
		outstr := string(out)
		if outstr != tt.expected {
			t.Errorf(fmt.Sprintf("Script %q expected %q actual %q",
				tt.script, tt.expected, outstr))
		}
	}
}
