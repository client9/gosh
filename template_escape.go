package conduit

import "bytes"

// JSEscapeTemplate escapes javascript-with-golang-templates into
// "pure" javascript.  Then the output can work with standard JS tooling
func JSEscapeTemplate(bytesin []byte) []byte {
	tmp := bytes.Replace(bytesin, []byte("{{"), []byte("/*{{"), -1)
	tmp = bytes.Replace(tmp, []byte("}}"), []byte("}}*/"), -1)
	return tmp
}

// JSUnescapeTemplate unescapes the above
func JSUnescapeTemplate(bytesin []byte) []byte {
	tmp := bytes.Replace(bytesin, []byte("/*{{"), []byte("{{"), -1)
	tmp = bytes.Replace(tmp, []byte("}}*/"), []byte("}}"), -1)
	return tmp
}

// HTMLEscapeTemplate escapes HTML-with-golang-templates into pure HTML
func HTMLEscapeTemplate(bytesin []byte) []byte {
	tmp := bytes.Replace(bytesin, []byte("{{"), []byte("<!--{{"), -1)
	tmp = bytes.Replace(bytesin, []byte("}}"), []byte("}}-->"), -1)
	return tmp
}

// HTMLUnescapeTemplate  Undoes the above
func HTMLUnescapeTemplate(bytesin []byte) []byte {
	tmp := bytes.Replace(bytesin, []byte("<!--{{"), []byte("{{"), -1)
	tmp = bytes.Replace(bytesin, []byte("}}-->"), []byte("}}"), -1)
	return tmp
}
