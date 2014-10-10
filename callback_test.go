package conduit

import (
	"bytes"
	"io/ioutil"
	"testing"
)

// MapAny is a string-to-any mapping
type MapAny map[string]interface{}

// Set sets an entry in the map
func (m MapAny) Set(key string, value interface{}) MapAny {
	m[key] = value
	return m
}

// Get returns an entry in the map
func (m MapAny) Get(key string) interface{} {
	return m[key]
}

func TestCallback(t *testing.T) {
	const script1 = `
.Set "foo" "bar"
`
	m := MapAny{}
	tpl, err := Parse("test1", script1)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(ioutil.Discard, m)
	if err != nil {
		panic(err)
	}
	if m["foo"] != "bar" {
		t.Errorf("Callback did not work")
	}

	const script2 = `
.foo
`
	tpl, err = Parse("test1", script2)
	if err != nil {
		panic(err)
	}
	out := bytes.Buffer{}
	err = tpl.Execute(&out, m)
	if err != nil {
		panic(err)
	}
	if out.String() != "bar" {
		t.Errorf("Got out %q", out.String())
	}
}

func TestCallback2(t *testing.T) {
	const script1 = `
$m := NewRegexpMatcher (array "foo" "bar" "abc")
$.Set "foo" $m
`

	const script2 = `
"abc" | ($.Get "foo").MatchOne 
`
	m := MapAny{}

	tpl := New("test2a").Funcs(MatchersFuncMap)
	tpl, err := tpl.Parse(TransformTemplate([]byte(script1)))
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(ioutil.Discard, m)
	if err != nil {
		panic(err)
	}

	tpl2 := New("test2b").Funcs(MatchersFuncMap)
	tpl2, err = tpl2.Parse(TransformTemplate([]byte(script2)))
	out := bytes.Buffer{}
	err = tpl2.Execute(&out, m)
	if err != nil {
		panic(err)
	}
	outstr := out.String()
	if outstr != "true" {
		t.Errorf("got %s", outstr)
	}
}

/*
func TestCallback3(t *testing.T) {
	const script1 = ` $.Foo = "bar"`
	tpl := New("test3").Funcs(MatchersFuncMap)
	tpl, err := tpl.Parse(TransformTemplate([]byte(script1)))
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	err = tpl.Execute(ioutil.Discard, m)
	if err != nil {
		panic(err)
	}
	if m["Foo"] != "bar" {
		t.Errorf("Setting directly did not work")
	}
}
*/
