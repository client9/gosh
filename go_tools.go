package gosh

// GoFmt is wrapper to "gofmt"
func GoFmt(args ...interface{}) ([]byte, error) {
	return StandardCLI("gofmt", nil, args)
}

// GoImports is wrapper to "goimports"
func GoImports(args ...interface{}) ([]byte, error) {
	return StandardCLI("goimports", nil, args)
}

// GoLint is a wrapper to "golint"
func GoLint(args ...interface{}) ([]byte, error) {
	return StandardCLI("golint", nil, args)
}

// GoVet is a wrapper to "go vet"
func GoVet(args ...interface{}) ([]byte, error) {
	return StandardCLI("go", []string{"vet"}, args)
}

// Git is a wrapper to "git"
func Git(args ...interface{}) ([]byte, error) {
	return StandardCLI("git", nil, args)
}
