deps:
	go get -d -t ./...

test: deps
	go test -v

lint:
	go vet
	golint -set_exit_status
