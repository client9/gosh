package conduit

import (
	"bytes"
)

func LineCount(bytesin []byte) int {
	count := 0
	for len(bytesin) > 0 {
		idx := bytes.IndexByte(bytesin, '\n')
		if idx == -1 {
			return count
		}
		count++
		bytesin = bytesin[idx+1:]
	}
	return count
}
