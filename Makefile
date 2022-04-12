postgres:
	docker run --name postgres-latest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres 
createdb: 
	docker exec -it postgres-latest createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres-latest dropdb --username=root simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up  
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down  
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate 
test: 
	go test -v -cover ./...
server:
	go run main.go
mockgen:
	mockgen -package mockdb -destination db/mock/store.go  9bany/simple_bank/db/sqlc Store 
.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mockgen migrateup1 migratedown1
