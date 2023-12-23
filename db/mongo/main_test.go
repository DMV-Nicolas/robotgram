package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/DMV-Nicolas/robotgram/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var testQueries Querier
var testCtx context.Context

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../.")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	testCtx = context.TODO()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort)
	client, err := mongo.Connect(testCtx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	db := client.Database(config.DBName)
	testQueries = NewQuerier(db)

	os.Exit(m.Run())
}
