package main

import (
	"context"
	"log"

	"github.com/DMV-Nicolas/tinygram/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DBSource))
	db := client.Database("tinygram")

	db.Drop(context.TODO())
}
