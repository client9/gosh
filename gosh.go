package gosh

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

// Transforms a input into a Go-Template
//
// # foo
// range glob "*"
//   println .
// end
//
// {{ range glob "*" }}{{ println . }}{{ end }}
//

// DefaultFuncMap is the default function map for the template engine
//
// Fundamental Types: strings, []strings
//
var DefaultFuncMap = template.FuncMap{
	"array": Array,
	"take":  ArrayTake,
	"drop":  ArrayDrop,
	// TODO sort, revert
	"assert":          Assert,
	"contains":        stringContains,
	"debug":           Debug,
	"discard":         DevNull,
	"grep":            Grep,
	"integer":         Integer,
	"hassuffix":       stringHasSuffix,
	"hasprefix":       stringHasPrefix,
	"join":            Join,
	"linecount":       LineCount,
	"lastindex":       stringLastIndex,
	"lower":           stringLower,
	"regexpcompile":   RegexpCompile,
	"regexpmatch":     RegexpMatch,
	"regexpquotemeta": RegexpQuoteMeta,

	"pathmatch": PathMatch,
	"repeat":    stringRepeat,
	"slice":     stringSlice,
	"string":    String,
	"split":     Split,
	"trim":      stringTrim,
	"trimleft":  stringTrimLeft,
	"trimright": stringTrimRight,
	"trimspace": stringTrimSpace,
	"type":      Type,
	"upper":     stringUpper,
}

// ExternalFuncMap defines extra functions that interface with external processes and file systems
var ExternalFuncMap = template.FuncMap{
	"now":                  TimeNow,
	"gzip":                 Gzip,
	"gunzip":               Gunzip,
	"git":                  Git,
	"glob":                 Glob,
	"open":                 Open,
	"mtime":                FileModTime,
	"gofmt":                GoFmt,
	"goimports":            GoImports,
	"golint":               GoLint,
	"govet":                GoVet,
	"stat":                 FileStat,
	"jsescapetemplate":     JSEscapeTemplate,
	"jsunescapetemplate":   JSUnescapeTemplate,
	"htmlescapetemplate":   HTMLEscapeTemplate,
	"htmlunescapetemplate": HTMLUnescapeTemplate,
	"jshint":               JSHint,
	"jsbeautify":           JSBeautify,
	"jsuglify":             JSUglify,
	"write":                Write,
}

// TransformTemplate turns a input string into a go template
func TransformTemplate(stdin []byte) string {
	out := bytes.Buffer{}
	lines := bytes.Split(stdin, []byte("\n"))
	continuation := false
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		if !continuation {
			out.WriteString("{{ ")
		}

		if line[len(line)-1] == '\\' {
			continuation = true
			line[len(line)-1] = ' '
			line = bytes.TrimSpace(line)
			line = append(line, ' ')
		} else {
			continuation = false
		}

		out.Write(line)

		if !continuation {
			out.WriteString(" }}")
		}
	}
	return out.String()
}

// New creates a new template with default functions defined
func New(name string) *template.Template {
	return template.New(name).Funcs(DefaultFuncMap)
}

// Execute -- hack
func Execute(name string, script string, data interface{}) ([]byte, error) {
	t := New(name)
	t, err := t.Parse(TransformTemplate([]byte(script)))
	if err != nil {
		return nil, err
	}
	out := bytes.Buffer{}
	err = t.Execute(&out, data)
	return out.Bytes(), err
}

// ExecuteExtended hack
func ExecuteExtended(name string, script string, data interface{}) ([]byte, error) {
	t := New(name)
	t = t.Funcs(ExternalFuncMap)
	t, err := t.Parse(TransformTemplate([]byte(script)))
	if err != nil {
		return nil, err
	}
	out := bytes.Buffer{}
	err = t.Execute(&out, data)
	return out.Bytes(), err
}

// Parse is a hack
func Parse(name string, script string) (*template.Template, error) {
	t := New(name)
	t, err := t.Parse(TransformTemplate([]byte(script)))
	if err != nil {
		return nil, err
	}
	return t, err
}

// MustExecute is a hack
func MustExecute(name string, script string, data interface{}) ([]byte, error) {
	out, err := Execute(name, script, data)
	if err != nil {
		os.Stdout.Write(out)
		os.Stdout.Sync()
		fmt.Fprintf(os.Stderr, "\n%v\n", err)
		os.Exit(1)
	}
	return out, nil
}
