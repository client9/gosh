package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/client9/conduit"
)

func main() {
	var raw []byte
	var err error
	var name string

	readFromString := flag.String("c", "", "Read from string")
	flag.Parse()
	args := flag.Args()
	if len(*readFromString) != 0 {
		name = "cli"
		raw = []byte(*readFromString)
	} else if len(args) == 1 {
		name = args[0]
		raw, err = ioutil.ReadFile(args[0])
		if err != nil {
			log.Fatalf("Unable to read %q: %s", args[0], err)
		}
	} else {
		name = "stdin"
		raw, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Unable to read stdin: %s", err)
		}
	}

	// TODO -- pass in CLI arguments

	t := template.New(name)
	t = t.Funcs(conduit.DefaultFuncMap)
	t = t.Funcs(conduit.ExternalFuncMap)
	t, err = t.Parse(conduit.TransformTemplate(raw))
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n%v\n", err)
		os.Exit(1)
	}
}
