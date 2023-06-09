package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
)

func OpenConnection() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	Database = client.Database("contacts")
	Collection = Database.Collection("users")

	log.Println("MongoDB connected")
}

func CloseConnection() {
	err := Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Mongo disconnected")
}

func GetConnection() *mongo.Client {
	if Client != nil {
		return Client
	}
	OpenConnection()
	return Client
}
