package conduit

import (
	"io/ioutil"
	"testing"
)

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
}
