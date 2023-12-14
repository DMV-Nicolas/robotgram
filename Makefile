test:
	go clean -testcache
	go test -v --cover ./...
server:
	go run main.go
.PHONY: test server