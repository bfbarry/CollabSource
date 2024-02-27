package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/bfbarry/CollabSource/back-end/mongoClient"
	"github.com/bfbarry/CollabSource/back-end/model"
)

type Repository struct {
	mongoClient *mongo.Database
}

var MongoRepository *Repository

func buildRepository(){
	mongoClient.InitilizeMongoClient()
	// defer mongoClient.CloseMongoClient()
	MongoRepository = &Repository{mongoClient: mongoClient.GetMongoClient()}
}

func GetMongoRepository() *Repository{
	if MongoRepository != nil {
		log.Println("Using existing repository object")
		return MongoRepository
	} else {
		log.Println("Creating new repository object")
		buildRepository()
		return MongoRepository
	}
	
}

func (self *Repository) Insert(collection string, object model.Model) {
	
	res, err := self.mongoClient.Collection(collection).InsertOne(context.TODO(), bson.D{{"name", "Alice"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
}

func (self *Repository) FindByID(collection string, id string){
	
	res, err := self.mongoClient.Collection(collection).InsertOne(context.TODO(), bson.D{{"name", "Alice"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
}

func (self *Repository) FindByQuery(collection string, query bson.M){
	
	cursor, err := self.mongoClient.Collection(collection).Find(context.TODO(),query)
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var result bson.M // You can use a custom struct type instead of bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	// fmt.Printf("inserted document with ID %v\n", res)
}

func (self *Repository) GetAll(collection string){
	
	cursor, err := self.mongoClient.Collection(collection).Find(context.TODO(),bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var result bson.M // You can use a custom struct type instead of bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	// fmt.Printf("inserted document with ID %v\n", res)
}

