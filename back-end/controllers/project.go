package controllers

import (
	"encoding/json"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"fmt"
)
// TODO move to models
type ProjectEnv struct {
	Coll  *mongo.Collection
}

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category 	string   `json:"category"`
	Tags        []string `json:"tags"`
	// DateCreated string
	// Members []string
	// Location    string   `json:"location"`
}
func (env *ProjectEnv) GetProjectByID(id string) []byte {
	var result Project
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objId}
	err = env.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	jsonResponse, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	return jsonResponse

}

func (env *ProjectEnv) GetProjectsByFilter(filterField string) []byte {
	// TODO: filter should be struct like Project struct
	var results []Project
	findOptions := options.Find()
	findOptions.SetLimit(20) // TODO: paginate properly
	filter := bson.M{"category": filterField}
	cursor, err := env.Coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(fmt.Sprintf("error finding projects: %s", err))
	}

	for cursor.Next(context.TODO())  {
		var elem Project
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	jsonResponse, err := json.Marshal(results)
	if err != nil {
		// TODO: handle error properly
		// http.Error(w, err.Error(), http.StatusInternalServerError)\
		log.Fatal(err)
		// return
	}
	return jsonResponse
}

func (env *ProjectEnv) CreateProject(p Project) []byte {
	// TODO: abstract away db
	env.Coll.InsertOne(context.TODO(), p)
	return []byte("success")
}
