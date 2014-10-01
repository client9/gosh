package conduit

// DevNull maps any input to empty string
// Useful in validators
func DevNull(data interface{}) string {
	return ""
}
