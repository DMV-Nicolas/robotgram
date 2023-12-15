package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/DMV-Nicolas/tinygram/util"
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
	client, err := mongo.Connect(testCtx, options.Client().ApplyURI(config.DBSource))
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	db := client.Database("tinygram")
	testQueries = New(db)

	os.Exit(m.Run())
}
