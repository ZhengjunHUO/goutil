.PHONY: test
test: 
	go test ./... -covermode=count -coverprofile=coverage.out
