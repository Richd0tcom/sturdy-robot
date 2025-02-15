package main

import (
	"context"
	"fmt"
	"log"

	api "github.com/Richd0tcom/sturdy-robot/internal/handlers"
	"github.com/Richd0tcom/sturdy-robot/internal/router"
	"github.com/jackc/pgx/v5/pgxpool"

	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"

	//utils
	"github.com/Richd0tcom/sturdy-robot/internal/config"
	// _ "github.com/lib/pq"
)

func main() {

	
	config, err:= config.LoadConfig(".env")
	if err != nil {
		fmt.Println("could not read config file")
	}

	// conn, err := sql.Open(config.DbDriver, config.DbUri)

	connPool, err:= pgxpool.New(context.Background(), config.DbUri)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store:= db.NewStore(connPool)
	server:= api.NewServer(store)

	router.SetupRouter(server)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server",err)
	}
}