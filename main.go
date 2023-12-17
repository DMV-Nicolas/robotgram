package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DMV-Nicolas/tinygram/api"
	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/DMV-Nicolas/tinygram/util"
	_ "github.com/golang/mock/mockgen/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// load config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect to database
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	// create an object queries for the database functions
	queries := db.NewQuerier(client.Database(config.DBName))

	// create server
	server, err := api.NewServer(config, queries)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	// start server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
