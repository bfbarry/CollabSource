package mongoClient

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"fmt"
	"os"
)

var mongoClient *mongo.Client
var dbName string

func init() {
	uri := os.Getenv("MONGODB_URI")
	dbName = os.Getenv("MONGODB_DB")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println("error connecting to mongodb, will exit program")
		panic(err)
	}
	
	log.Printf("Successfully connected to mongo at %s",uri)

	var result bson.M
	if err := GetMongoClient().RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	mongoClient = client

}

func CloseMongoClient() {
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		log.Println("error closing mongodb connection")
		log.Fatal(err)
	}
}

func GetMongoClient() *mongo.Database{
	return mongoClient.Database(dbName)
}

