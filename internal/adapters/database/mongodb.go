package database

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/RiteshDevOpsEngineer/ecom/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoOnce   sync.Once
	mongoClient *mongo.Client
)

func InitializeMongoDB() (*mongo.Client, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	clientOptions := options.Client().ApplyURI(cfg.MongoDB.ConnectionURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// log.Println("Connected to MongoDB!")
	return client, nil
}

func GetMongoClient() *mongo.Client {
	mongoOnce.Do(func() {
		var err error
		mongoClient, err = InitializeMongoDB()
		if err != nil {
			log.Fatal("Failed to initialize MongoDB client:", err)
		}
	})
	return mongoClient
}
