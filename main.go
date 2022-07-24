package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/syafrin-ibrahim/mybank/api"
	db "github.com/syafrin-ibrahim/mybank/db/sqlc"
	"github.com/syafrin-ibrahim/mybank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
