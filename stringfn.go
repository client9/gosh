package conduit

import (
	"errors"
	"fmt"
	"strings"
)

func stringSlice(args ...interface{}) (string, error) {
	var fromIdx, untilIdx int
	var source string
	var err error
	switch len(args) {
	case 2:
		fromIdx, err = Integer(args[0])
		if err != nil {
			return "", err
		}
		source = String(args[1])
		untilIdx = len(source)
	case 3:
		fromIdx, err = Integer(args[0])
		if err != nil {
			return "", err
		}
		untilIdx, err = Integer(args[1])
		if err != nil {
			return "", err
		}
		source = String(args[2])
	default:
		return "", errors.New("Requires 2 or 3 arguments")
	}
	// -1 is ok since it's assumed its from an index
	if fromIdx == -1 || untilIdx == -1 {
		return "", nil
	}
	if fromIdx > untilIdx {
		return "", fmt.Errorf("Slice [%d:%d] is invalid", fromIdx, untilIdx)
	}
	if untilIdx > len(source) {
		return "", fmt.Errorf("Slice is greater than length")
	}
	return source[fromIdx:untilIdx], nil
}

func stringContains(sep string, source string) bool {
	return strings.Index(source, sep) != -1
}

func stringIndex(sep string, source string) int {
	return strings.Index(source, sep)
}

func stringLastIndex(sep string, source string) int {
	return strings.LastIndex(source, sep)
}

func stringRepeat(count int, source string) string {
	return strings.Repeat(source, count)
}

func stringLower(source string) string {
	return strings.ToLower(source)
}
func stringLowerMulti(sources ...string) []string {
	for pos, val := range sources {
		sources[pos] = strings.ToLower(val)
	}
	return sources
}

func stringUpper(source string) string {
	return strings.ToUpper(source)
}

func stringTrimSpace(source string) string {
	return strings.TrimSpace(source)
}

func stringHasPrefix(prefix string, source string) bool {
	return strings.HasPrefix(source, prefix)
}
func stringHasSuffix(suffix string, source string) bool {
	return strings.HasSuffix(source, suffix)
}

func stringTrimRight(cutset string, source string) string {
	return strings.TrimRight(source, cutset)
}

func stringTrimLeft(cutset string, source string) string {
	return strings.TrimLeft(source, cutset)
}

func stringTrimPrefix(prefix string, source string) string {
	return strings.TrimPrefix(source, prefix)
}

func stringTrimSuffix(suffix string, source string) string {
	return strings.TrimSuffix(source, suffix)
}

func stringTrim(cutset string, source string) string {
	return strings.Trim(source, cutset)
}
