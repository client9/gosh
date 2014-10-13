package gosh

import (
	"testing"
)

func TestTransformations(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{
			in:  "line",
			out: "{{ line }}",
		},
		{
			in:  "line1\nline2",
			out: "{{ line1 }}{{ line2 }}",
		},
		{
			in:  "line1a\\\nline1b\nline2",
			out: "{{ line1a line1b }}{{ line2 }}",
		},
	}

	for _, tt := range tests {
		actual := TransformTemplate([]byte(tt.in))
		if tt.out != actual {
			t.Errorf("got %s expected %s", actual, tt.out)
		}
	}
}
