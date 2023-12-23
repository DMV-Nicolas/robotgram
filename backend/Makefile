mongo:
	docker run -d -p27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=secret --name mongo mongo:latest
server:
	go run main.go
test:
	go clean -testcache
	go test -v --cover ./...
mock:
	mockgen -package mockdb -destination db/mock/queries.go github.com/DMV-Nicolas/robotgram/db/mongo Querier
dropdb:
	go run ./commands/dropdb/main.go
.PHONY: docker test server mock
