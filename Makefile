postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=tsarka -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=tsarka --owner=tsarka tsarka_test

dropdb:
	docker exec -it postgres15 dropdb tsarka_test

migrateup:
	migrate -path db/migration -database "postgresql://tsarka:secret@localhost:5432/tsarka_test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://tsarka:secret@localhost:5432/tsarka_test?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run cmd/main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server

