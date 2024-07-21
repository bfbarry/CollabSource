package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)


func (self *Repository) DocumentExists(coll string, id primitive.ObjectID) (bool, error) {
	filter := bson.M{"_id": id}
	count, err := self.getCollection(coll).CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} 
	return false, nil
}

