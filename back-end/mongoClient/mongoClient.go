package mongoClient

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"fmt"
)

const uri = "mongodb://localhost:27017"

var mongoClient *mongo.Client;

func InitilizeMongoClient() {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println("error connecting to mongodb, will exit program")
		panic(err)
	}
	
	log.Println("Successfully connected to mongo at %s",uri)

	var result bson.M
	if err := client.Database("test").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	mongoClient = client

}

func CloseMongoClient() {
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		log.Println("error closing mongodb connection")
		panic(err)
	}
}

func GetMongoClient() *mongo.Database{
	return mongoClient.Database("test")
}

