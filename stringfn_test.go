package conduit

import (
	"fmt"
	"testing"
)

func TestLower(t *testing.T) {
	var casetests = []struct {
		script   string
		expected string
	}{
		{` lower $`, "foo.png"},
		{` $ | lastindex "." `, "3"},
		{` slice (lastindex "." $) $ `, ".PNG"},
		{` slice (lastindex "." $) $ | lower | contains ".png" `, "true"},
		//	{` $ | slice (lastindexof "." $ ) `, ".PNG" },
		//	{` lowerm "FOO" "BAR" `, "[foo bar]"},
		//	{` split " " "FOO BAR" | lower `, "[foo bar]"},
	}

	for _, tt := range casetests {

		out, err := Execute("testlower", tt.script, "FOO.PNG")
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
