package conduit

import (
	"io/ioutil"
	"testing"

	"github.com/cloudflare/ahocorasick"
)

const script1 = `
$ua := .
if array  "sqlmap" "nmap" "nikto" "paros" "zmeu" "morfeus" "havij" "netsparker" "w3af.sourceforge.net" "Mozilla/4.0 (compatible; Synapse)" | substring_match $ua
      printf .
end 
`
const ua = "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2049.0 Safari/537.36"

type AhoCorasick struct {
	matcher *ahocorasick.Matcher
}

// NewAhoCorasick is a constructure
func NewAhoCorasick(substrings []string) *AhoCorasick {
	return &AhoCorasick{
		matcher: ahocorasick.NewStringMatcher(substrings),
	}
}

// MatchOne matches the first substring found
//  -1 if no match, else index of first match
func (m *AhoCorasick) MatchOne(s string) int {
	indexes := m.matcher.Match([]byte(s))
	if len(indexes) == 0 {
		return -1
	}
	return indexes[0]
}

func BenchmarkSearch1(b *testing.B) {
	t, err := Parse("test1", script1)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		err := t.Execute(ioutil.Discard, ua)
		if err != nil {
			panic(err)
		}
	}
}

const script2 = `
if substring_match .UserAgent .Substring
	printf .
end
`

func BenchmarkSearch2(b *testing.B) {
	ud := make(map[string]interface{})
	ud["UserAgent"] = ua
	ud["Substring"] = []string{"sqlmap", "nmap", "nikto", "paros", "zmeu", "morfeus", "havij", "netsparker", "w3af.sourceforge.net", "Mozilla/4.0 (compatible; Synapse)"}
	t, err := Parse("test2", script2)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		err := t.Execute(ioutil.Discard, ud)
		if err != nil {
			panic(err)
		}
	}
}

const script3 = `
if .Aho.MatchOne .UserAgent
	println ".UserAgent	"
end
`

func BenchmarkSeach3(b *testing.B) {
	ud := make(map[string]interface{})
	substrings := []string{"sqlmap", "nmap", "nikto", "paros", "zmeu", "morfeus", "havij", "netsparker", "w3af.sourceforge.net", "Mozilla/4.0 (compatible; Synapse)"}
	ud["UserAgent"] = ua
	ud["Aho"] = NewAhoCorasick(substrings)
	t, err := Parse("test3", script3)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		err := t.Execute(ioutil.Discard, ud)
		if err != nil {
			panic(err)
		}
	}
}
