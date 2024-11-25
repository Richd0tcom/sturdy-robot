# Easily setup the development environment by running make <command name>
postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Madara123 -d postgres:15-alpine
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root numer_invoice
dropdb:
	docker exec -it postgres15 dropdb --username=root  numer_invoice

createmig:
	migrate create -ext sql -dir internal/db/migrations -seq init
migrateup:
	migrate -path internal/db/migrations -database "postgresql://root:Madara123@localhost:5432/numer_invoice?sslmode=disable" -verbose up
migratedown:
	migrate -path internal/db/migrations -database "postgresql://root:Madara123@localhost:5432/numer_invoice?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
mock:
# the --build_flags=--mod=mod may be removed in the future.should only be used if you encounter errors on th first run.
	mockgen -package mockdb -destination db/mock/store.go github.com/Richd0tcom/sturdy-robot/db/sqlc Store

server:
	go run main.go

.PHONY: postgres dropdb createdb createmig migrateup migratedown test server