package main

import (
	"context"
	"log"

	"github.com/DMV-Nicolas/tinygram/api"
	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/DMV-Nicolas/tinygram/util"
	_ "github.com/golang/mock/mockgen/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DBSource))
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	queries := db.New(client.Database(config.DBName))

	server, err := api.NewServer(config, queries)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
