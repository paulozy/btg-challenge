package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewDatabaseConnection(dsn string) (*mongo.Client, func()) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(dsn),
	)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client, func() {
		fmt.Println("discconecting database")
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
}
