test:
	go clean -testcache
	go test -v --cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/DMV-Nicolas/tinygram/db/mongo Querier
dropdb:
	go run ./commands/dropdb/main.go
.PHONY: test server mock