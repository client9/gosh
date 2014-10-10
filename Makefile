
build:
	go install -v ./...
	go test -v ./...

test:
	go test -v ./...

benchmark:
	go test -test.benchmem -bench .

clean:
	go clean

