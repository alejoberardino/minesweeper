package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

// GetConnection - Retrieves a client to the DocumentDB
func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	clusterEndpoint := os.Getenv("MONGO_HOST")

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)
	log.Printf("Connecting to client %s", connectionURI)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	// We want the connection for this service to live forever in the background
	ctx, cancel := context.WithCancel(context.Background())

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}
