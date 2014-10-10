gosh
=======

Golang template-based "shell"

## What is it

TODO

## GOSH to go-lang template

TODO

## Example: Let's check Go source files with goimports


```
goimport "-l" "."
```

runs `goimport` on the current directory.  This is fine more many cases, but as an example will show how to select different files and process the output.

```
goimports "-l" (glob "*.go") | linecount | eq 0 | assert "goimports check failed"
```

* goimports (the external program) is called by computing some command line arguments with glob.
* linecount takes that an produces an integer
* eq is a standard golang template function, here is the linecount 0?
* assert, if false, it errors and halt execution.  If true does nothing.


That is "fail if the output of `goimport -l *.go` is empty".

NOTE: It is possible to have StandardCLI check the last argument for `*` and automatically do
a glob expland.  Then it would just be 'goimport "-l" "*go".  Im not sure if that is good
or bad idea.


However it might be simpler to understand it this.

```
glob "*.go" | goimports "-l" | linecount | eq 0 | assert "goimports check failed"
```

* glob produces a list ([]string) or file names
* goimports uses that to produce a full call.
* the rest if the same


A more friendly version would be this

```
with glob "*.go" | goimports "-l" | split
        println "Correct the following files: "
        range .
                printf "goimports -w %s\n" .
        end
        assert false
end
```


Another way using stdin asking for the diff ouptut.


```
#!/usr/bin/env gosh
range $fname := glob "*.go"
    open . | goimports "-d" | linecount | eq 0 | assert ( printf "%s : goimport failed" $fname )
end
```

Checks that each file produces no diff output from  `goimports` (i.e. nothing changed), else asserts.  


A similar example from above, but the main pipeline operates on the data file itself.

```
#!/usr/bin/env gosh
range $fname := glob "*.go"
    with open $fname
        . | goimports "-d" | linecount |  eq 0 | assert ( printf "failed: goimport -d %q" $fname )
    end
end
```

Re-using code

```
define "goimport-check"
    . | goimports "-l" | linecount | eq 0 | assert "goimports check failed"	
end
	
template "goimport-check" (glob "*.go")
```

```
#!/usr/bin/env gosh

define "goimport-check"
    with . | goimports "-l" | split
        println "Correct the following files: "
        range .
            printf "goimports -w %s\n" .
        end
        assert false
    end
end

with glob "*.go"
    template "goimport-check" .
end
```
