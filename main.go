package main

import (
	"9bany/simple_bank/api"
	db "9bany/simple_bank/db/sqlc"
	"9bany/simple_bank/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("Can not load config file: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NerServer(store)
	server.Start(config.ServerAddress)
}
