package mongoClient

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/joho/godotenv"
	"log"
	"fmt"
	"os"
)

var mongoClient *mongo.Client
var dbName string

func init() {
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
	uri := os.Getenv("MONGODB_URI")
	dbName = os.Getenv("MONGODB_DB")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("error connecting to mongodb, will exit program")
	}
	
	log.Printf("Successfully connected to mongo at %s",uri)
	mongoClient = client

	var result bson.M
	if err := GetMongoDb().RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatalf("Error in ping %s", err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func ClosemongoClient() {
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		log.Println("error closing mongodb connection")
		log.Fatal(err)
	}
}

func GetMongoDb() *mongo.Database{
	return mongoClient.Database(dbName)
}

