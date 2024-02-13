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
type ProjectEnv struct {
	Coll  *mongo.Collection
}

// TODO move to models
type Project struct {
	Name        string   `json:"name"        bson:"name,omitempty"`
	Description string   `json:"description" bson:"description,omitempty"`
	Category 	string   `json:"category"    bson:"category,omitempty"`
	Tags        []string `json:"tags"        bson:"tags,omitempty"`
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

func (env *ProjectEnv) UpdateProject(id string, p Project) []byte {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	res, err := env.Coll.UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": p})
	if err != nil {	
		log.Fatal(err)
	}
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		// TODO: handle error properly
		// http.Error(w, err.Error(), http.StatusInternalServerError)\
		log.Fatal(err)
		// return
	}
	return jsonResponse
}

func (env *ProjectEnv) DeleteProject(deleteMode DeleteMode, id string) []byte {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	var err2 error
	switch deleteMode {
		case SoftDelete:
			_, err2 = env.Coll.UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": bson.M{"deleted": true}})
		case HardDelete:
			_, err2 = env.Coll.DeleteOne(context.TODO(), bson.M{"_id": objId})
		}
	if err2 != nil {	
		log.Fatal(err)
	}

	return []byte("success")
}