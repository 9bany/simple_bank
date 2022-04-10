package main

import (
	"9bany/simple_bank/api"
	db "9bany/simple_bank/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	address  = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NerServer(store)
	server.Start(address)
}
