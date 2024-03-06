package repository

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/bfbarry/CollabSource/back-end/model"
	"log"
	"io"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func streamToObj(coll string, reqBody *io.ReadCloser) (model.Model, error) {
	obj := model.GetModelFromName(coll)
	err := json.NewDecoder(*reqBody).Decode(&obj)
	if err != nil { //TODO: 400
		log.Println("streamToObj - Error unmarshaling JSON: ", err)
		return nil, err
	}
	return obj, nil
}

func streamToBsonM(coll string, reqBody *io.ReadCloser) (bson.M, error) {
	obj, err := streamToObj(coll, reqBody)
	if err != nil {
		return nil, err
	}
	bsonBytes, err := bson.Marshal(obj)
	if err != nil {
		log.Fatalf("Error marshaling BSON: %v", err)
	}
	var result bson.M
	if err := bson.Unmarshal(bsonBytes, &result); err != nil {
		log.Fatalf("Error unmarshaling BSON: %v", err)
	}
	return result, nil
}

func cursorToSlice(cursor *mongo.Cursor, coll string) ([]model.Model, error) {
	results := []model.Model{}
	for cursor.Next(context.Background()) {
		result := model.GetModelFromName(coll)
		
		err := cursor.Decode(result)
		if err != nil {
			log.Println("Error decoding cursor: ", err)
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}