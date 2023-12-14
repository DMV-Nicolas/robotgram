createdb:
	docker exec -it tinygram createdb tinygramdata
dropdb:
	docker exec -it tinygram dropdb tinygramdata
migrateup:
	migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/tinygramdata?sslmode=disable" -verbose up 
migratedown:
	migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/tinygramdata?sslmode=disable" -verbose down
test:
	go clean -testcache
	go test -v --cover ./...
server:
	go run main.go
.PHONY: createdb dropdb migrateup migratedown test server