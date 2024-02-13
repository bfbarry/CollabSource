package connections

import (
	"os"
	"fmt"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// thanks https://www.alexedwards.net/blog/organising-database-access
// used for dependency injection
type Env struct {
	DB *mongo.Database
}

func InitDB() (*mongo.Client, *mongo.Database) {
	DBNAME := "CollabSourceDev0"
	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		fmt.Println("MONGODB_URI not set")
		os.Exit(1)
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("error connecting to mongodb in InitDB")
		panic(err)
	}

	db := client.Database(DBNAME)
	return client, db
}

func CloseDB(client *mongo.Client) func() {
	//useage: defer CloseDB(client)()
	return func() {
		log.Println("disconnecting from mongodb")
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(fmt.Sprintf("disconnected on defer: %s", err))
		}
	}
}