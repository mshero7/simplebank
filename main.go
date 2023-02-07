package main

import (
	"database/sql"
	"log"

	"github.com/mshero7/simplebank/api"
	db "github.com/mshero7/simplebank/db/sqlc"
	"github.com/mshero7/simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewSever(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start sesrver")
	}
}
