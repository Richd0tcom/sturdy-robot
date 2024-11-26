package main

import (
	"database/sql"
	"log"

	api "github.com/Richd0tcom/sturdy-robot/internal/handlers"

	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
	
	//utils
	"github.com/Richd0tcom/sturdy-robot/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	config, err:= config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not read configs ", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbUri)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store:= db.NewStore(conn)
	server:= api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server",err)
	}
}