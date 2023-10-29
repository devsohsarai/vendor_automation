package configs

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var (
// 	databaseName = DbMongo()
// )

func ConnectDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(EnvMongoURI())
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Ping the database to establish a connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(DbMongo()).Collection(collectionName)
	return collection
}

func DisconnectDB() {
	if err := DB.Disconnect(context.Background()); err != nil {
		panic(err.Error())
	}
	fmt.Println("Connection has disconnect!")
}
