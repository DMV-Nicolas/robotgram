// This command works for to drop the database
package main

import (
	"context"
	"fmt"
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

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	db := client.Database("tinygram")

	db.Drop(context.TODO())
}
