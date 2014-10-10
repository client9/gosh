package conduit

import (
	"os"
	"testing"
)

const script = `
$files := git "diff" "--cached" "--name-only" "--diff-filter=ACM" | pathmatch "*.go"

with $files | gofmt "-s" "-l" | split
        println "Correct the following files: "
        range .
                printf "gofmt -w -s %s\n" .
        end
        assert false
end

with $files | goimports "-l" | split
        println "Correct the following files: "
        range .
                printf "goimports -w %s\n" .
        end
        assert false
end

with $files | golint  | string
	print .
	assert false
end

`

// func Execute(name string, script string, data interface{}) ([]byte, error) {

func TestCodeFormat(t *testing.T) {
	out, err := ExecuteExtended("build_test", script, nil)
	if err != nil {
		os.Stderr.Write(out)
		t.Errorf(err.Error())
	}
}
