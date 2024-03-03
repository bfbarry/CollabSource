package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"log"
	"github.com/bfbarry/CollabSource/back-end/model"
)

func cursorToSlice(cursor *mongo.Cursor, coll string) []model.Model {
	var results []model.Model
	for cursor.Next(context.TODO()) {
		elem := model.GetModelFromName(coll)
		err := cursor.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	return results
}