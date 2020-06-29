
make: fmt

fmt:
	go fmt ./...

lint:
	golint ./...