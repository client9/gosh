package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/client9/conduit"
	"io/ioutil"
	"log"
	"text/template"
)

/*

gosh 'template' files

*/

func main() {
	inplace := false
	flag.BoolVar(&inplace, "inplace", false, "overwrite input files")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatalf("Usages: gosh 'pipeline' files....")
	}

	cmds := "{{ . | " + args[0] + " | string }}"
	//cmds := "{{ . | " + args[0] + " }}"
	funcMap := template.FuncMap{
		"jsescapetemplate":     conduit.JSEscapeTemplate,
		"jsunescapetemplate":   conduit.JSUnescapeTemplate,
		"htmlescapetemplate":   conduit.HTMLEscapeTemplate,
		"htmlunescapetemplate": conduit.HTMLUnescapeTemplate,
		"linecount":            conduit.LineCount,
		"gzip":                 conduit.Gzip,
		"gunzip":               conduit.Gunzip,
		"string":               conduit.String,
		"jshint":               conduit.JSHint,
		"jsbeautify":           conduit.JSBeautify,
		"jsuglify":             conduit.JSUglify,
		"devnull":              conduit.DevNull,
	}

	t, err := template.New("cli").Funcs(funcMap).Parse(cmds)
	if err != nil {
		log.Fatalf("Unable to parse %s: %s", cmds, err)
	}

	// TOD use stdin
	files := args[1:]
	for _, fname := range files {
		fbytes, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatalf("Error reading %s: %s", fname, err)
		}
		out := bytes.Buffer{}
		err = t.Execute(&out, fbytes)
		if err != nil {
			log.Fatalf("%s: %v", fname, err)
			return
		}
		fmt.Printf("%s", out.String())
	}
}
