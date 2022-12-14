package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//this function creates a new mongodb connection
func New(cfg Config) (*mongo.Database, error) {
	opts := options.Client()
	opts.ApplyURI(cfg.URL)

	//create mongodb connection
	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, fmt.Errorf("db new client error:%w", err)
	}
	//connect to the mongo db
	{
		ctx, done := context.WithTimeout(context.Background(), cfg.ConnectionTimeout)

		defer done()

		if err := client.Connect(ctx); err != nil {
			return nil, fmt.Errorf("db connection error: %w", err)
		}
	}
	//ping the mongodb
	{
		ctx, done := context.WithTimeout(context.Background(), cfg.ConnectionTimeout)
		defer done()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			return nil, fmt.Errorf("db ping error: %w", err)
		}
	}
	return client.Database(cfg.Name), nil

}
