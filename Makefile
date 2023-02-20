postgres:
	docker run --name postgres -p 5432:5432 --network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
	
sqlc_mac:
	sqlc generate

test:
	go test -v -cover ./...
	
mock:
	mockgen -destination ./db/mock/store.go -package mockdb github.com/mshero7/simplebank/db/sqlc Store

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test mockgen migratedown1