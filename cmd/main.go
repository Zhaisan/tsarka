package main

import (
	"database/sql"
	"github.com/Zhaisan/tsarka_test/api"
	db "github.com/Zhaisan/tsarka_test/db/sqlc"
	"github.com/Zhaisan/tsarka_test/db/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	database := db.New(conn)
	server:= api.NewServer(database)
	
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
}
