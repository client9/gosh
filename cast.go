package conduit

import (
	"fmt"
)

func String(data interface{}) string {
	switch data.(type) {
	case string:
		bits, _ := data.(string)
		return bits
	case []byte:
		bits, _ := data.([]byte)
		return string(bits)
	default:
		return fmt.Sprintf("%v", data)
	}
}
