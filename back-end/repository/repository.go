package repository

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/bfbarry/CollabSource/back-end/model"
	"github.com/bfbarry/CollabSource/back-end/mongoClient"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	mongoClient *mongo.Database
}

var MongoRepository *Repository

func init() {
	// defer mongoClient.CloseMongoClient()
	MongoRepository = &Repository{mongoClient: mongoClient.GetMongoClient()}
}

func GetMongoRepository() *Repository{
	log.Println("Using existing repository object")
	return MongoRepository
}

func (self *Repository) getCollection(coll string) *mongo.Collection {
	return self.mongoClient.Collection(coll)
}

func (self *Repository) Insert(coll string, object model.Model) ([]byte, error) {
	res, err := self.getCollection(coll).InsertOne(context.TODO(), object)
	if err != nil {
		var writeErr mongo.WriteException
		if errors.As(err, &writeErr) {
			log.Println("WriteException in Insert")
			for _, we := range writeErr.WriteErrors {
				log.Println(we)
			}
		} else {
			log.Println("Error in Insert", err)
		}
		return nil, errors.New("error inserting document")
	}
	log.Printf("inserted document with ID %v\n", res.InsertedID)
	return []byte("success"), nil
}

func (self *Repository) FindByID(coll string, id string) (model.Model, error) {
	obj := model.GetModelFromName(coll)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objId}
	err = self.getCollection(coll).FindOne(context.TODO(), filter).Decode(&obj)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error finding document")
	}
	return obj, nil
}

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

//filter e.g, bson.M{"category": filterField}
// TODO: filter should be struct like Project struct
func (self *Repository) FindByQuery(coll string, filter bson.M) ([]model.Model, error){
	findOptions := options.Find()
	findOptions.SetLimit(20) // TODO: paginate properly
	cursor, err := self.getCollection(coll).Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error finding documents")
	}

	results := cursorToSlice(cursor, coll)
	return results, nil
}

func (self *Repository) FindAll(coll string) []byte {
	findOptions := options.Find()
	findOptions.SetLimit(20) // TODO: paginate properly
	cursor, err := self.getCollection(coll).Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	results := cursorToSlice(cursor, coll)
	jsonResponse, err := json.Marshal(results)
	if err != nil {
		// TODO: handle error properly
		// http.Error(w, err.Error(), http.StatusInternalServerError)\
		log.Fatal(err)
		// return
	}
	return jsonResponse
}

func (self *Repository) Update(coll string, obj model.Model, id string) ([]byte, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := self.getCollection(coll).UpdateOne(context.TODO(), 
												   bson.M{"_id": objId},
												   bson.M{"$set": obj})
	if err2 != nil {
		log.Println(err2)
		return nil, errors.New("error updating document")
	}
	log.Printf("updated document in %s, id: %s", coll, id)
	return []byte("success"), nil
}

func (self *Repository) Delete(coll string, deleteMode DeleteMode, id string) ([]byte, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	var del_err error
	switch deleteMode {
		case SoftDelete:
			_, del_err = self.getCollection(coll).UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": bson.M{"deleted": true}})
		case HardDelete:
			_, del_err = self.getCollection(coll).DeleteOne(context.TODO(), bson.M{"_id": objId})
		}
	if del_err != nil {	
		log.Println(del_err)
		return nil, errors.New("error deleting document")
	}

	return []byte("success"), nil
}