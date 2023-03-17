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

migrateup-aws:
	migrate -path db/migration -database "postgresql://postgres:12341234@simple-bank.cnkwfh9yit7c.ap-northeast-2.rds.amazonaws.com:5432/simple_bank" -verbose up

migratedown-aws:
	migrate -path db/migration -database "postgresql://postgres:12341234@simple-bank.cnkwfh9yit7c.ap-northeast-2.rds.amazonaws.com:5432/simple_bank" -verbose down

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

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt paths=source_relative \
    proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test mockgen migratedown1 proto evans