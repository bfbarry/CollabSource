package repository

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)


func (self *Repository) DocumentExists(coll string, id primitive.ObjectID) bool {
	filter := bson.M{"_id": id}
	count, err := self.getCollection(coll).CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		return true
	} 
	return false
}

