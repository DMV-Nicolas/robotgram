package db

import "go.mongodb.org/mongo-driver/mongo"

// NewQuerier creates a new querier
func NewQuerier(db *mongo.Database) Querier {
	return &Queries{db: db}
}

// Queries is an struct for to interact with the database
type Queries struct {
	db *mongo.Database
}
