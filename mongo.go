package vps_monitoring_notifier

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func pingMongo(mongoURI string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("error pinging MongoDB: %w", err)
	}

	return nil
}
