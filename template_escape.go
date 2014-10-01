package conduit

import "bytes"

func JSEscapeTemplate(bytesin []byte) []byte {
	tmp := bytes.Replace(bytesin, []byte("{{"), []byte("/*{{"), -1)
	tmp = bytes.Replace(tmp, []byte("}}"), []byte("}}*/"), -1)
	return tmp
}

func JSUnescapeTemplate(bytesin []byte) []byte {
	tmp := bytes.Replace(bytesin, []byte("/*{{"), []byte("{{"), -1)
	tmp = bytes.Replace(tmp, []byte("}}*/"), []byte("}}"), -1)
	return tmp
}

func HTMLEscapeTemplate(bytesin []byte) []byte {
	tmp := bytes.Replace(bytesin, []byte("{{"), []byte("<!--{{"), -1)
	tmp = bytes.Replace(bytesin, []byte("}}"), []byte("}}-->"), -1)
	return tmp
}

func HTMLUnescapeTemplate(bytesin []byte) []byte {
	tmp := bytes.Replace(bytesin, []byte("<!--{{"), []byte("{{"), -1)
	tmp = bytes.Replace(bytesin, []byte("}}-->"), []byte("}}"), -1)
	return tmp
}
