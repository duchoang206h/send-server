package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg *MongoInstance

func ConnectMongo(mongoURI, dbName string) error {
	// load config

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	db := client.Database(dbName)
	if err != nil {
		return err
	}
	mg = &MongoInstance{
		Client: client,
		Db:     db,
	}
	fmt.Println("mongodb connected")
	return nil
}

func GetMongo() *MongoInstance {
	return mg
}

func (mg *MongoInstance) Collection(name string) *mongo.Collection {
	collection := mg.Db.Collection(name)
	return collection
}

func (mg *MongoInstance) Disconnect(ctx context.Context) error {
	return mg.Client.Disconnect(ctx)
}
