package controllers

import (
	"net/http"
	"encoding/json"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"fmt"
)
// TODO move to models
type Project struct {
	_ID 			string	 `json:"_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	// DateCreated string
	// Members []string
	// Location    string   `json:"location"`
}

func (env *Env) GetSampleProjects(w http.ResponseWriter, r *http.Request) {

	var results []Project
	findOptions := options.Find()
	findOptions.SetLimit(20)
	cursor, err := env.DB.Collection("projects").Find(context.TODO(), bson.D{{}}, findOptions)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}