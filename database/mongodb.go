package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var mongoUrl = os.Getenv("DBConnectionStr")
var username = os.Getenv("Username")
var password = os.Getenv("Password")

func ConnectDB() (*mongo.Client, error) {
	credential := options.Credential{
		Username: username,
		Password: password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.
		Client().
		ApplyURI(mongoUrl).
		SetAuth(credential).
		SetHeartbeatInterval(5 * time.Second)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return client, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return client, err
	}

	return client, nil
}

func GetCollection(client *mongo.Client, name string, collectionName string) *mongo.Collection {
	return client.Database(name).Collection(collectionName)
}
