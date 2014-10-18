package gosh

import (
	"fmt"
	"testing"
)

func TestLower(t *testing.T) {
	var casetests = []struct {
		script   string
		input    string
		expected string
	}{
		{` lower $`, "FOO.PNG", "foo.png"},
		{` $ | lower`, "FOO.PNG", "foo.png"},
		{` $ | lower | upper`, "Foo.Png", "FOO.PNG"},
		{` $ | lower | repeat 2`, "FOO.PNG", "foo.pngfoo.png"},
		{` $ | lastindex "." `, "foo.png", "3"},
		{` slice (lastindex "." $) $ `, "foo.PNG", ".PNG"},
		{` slice (lastindex "." $) $ | lower | contains ".png" `, "foo.PNG", "true"},
		{` $ | trim "ax" `, "xaxfoo.pngxax", "foo.png"},
		{` $ | trimspace`, "  foo.png   ", "foo.png"},
		{` array "foo" "bar" "c" | drop 1 | join "-"`, "", "foo-bar"},
		//	{` $ | slice (lastindexof "." $ ) `, ".PNG" },
		//	{` lowerm "FOO" "BAR" `, "[foo bar]"},
		//	{` split " " "FOO BAR" | lower `, "[foo bar]"},
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
