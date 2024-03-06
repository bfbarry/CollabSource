package repository

import (
	"context"
	"errors"
	"log"
	"io"
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
	log.Println("initializing repository")
	MongoRepository = &Repository{mongoClient: mongoClient.GetMongoDb()}
}

func GetMongoRepository() *Repository{
	log.Println("Using existing repository object")
	return MongoRepository
}

func (self *Repository) getCollection(coll string) *mongo.Collection {
	return self.mongoClient.Collection(coll)
}

func (self *Repository) Insert(coll string, streamObj *io.ReadCloser) ([]byte, error) {
	obj, err := streamToObj(coll, streamObj)
	if err != nil { // TODO: 400
		return nil, err
	}
	res, err := self.getCollection(coll).InsertOne(context.TODO(), obj)
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
		return nil, err
	}
	log.Printf("inserted document with ID %v\n", res.InsertedID)
	return []byte("success"), nil
}

func (self *Repository) FindByID(coll string, id string) (model.Model, error) {
	obj := model.GetModelFromName(coll)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("error in FindByID %s", err)
		return nil, err
	}
	filter := bson.M{"_id": objId}
	err = self.getCollection(coll).FindOne(context.TODO(), filter).Decode(obj)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return obj, nil
}

//filter e.g, bson.M{"category": filterField}
//		bson.M{} for no filter
// TODO: set limit on pageSize
func (self *Repository) Find(coll string, streamFilterObj *io.ReadCloser, pageIndex int64, pageSize int64) ([]model.Model, error){
	filter, err := streamToBsonM(coll, streamFilterObj)
	if err != nil {
		return nil, err
	}
	findOptions := options.Find()
	skip := pageIndex * pageSize
	findOptions.SetSkip(skip)
	findOptions.SetLimit(pageSize)
	cursor, err := self.getCollection(coll).Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	results, err := cursorToSlice(cursor, coll)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return results, nil
}

func (self *Repository) Update(coll string, streamObj *io.ReadCloser, id string) ([]byte, error) {
	obj, err := streamToObj(coll, streamObj)
	if err != nil { // TODO: 400
		return nil, err
	}
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalf("Update %s", err)
	}
	_, err2 := self.getCollection(coll).UpdateOne(context.TODO(), 
												   bson.M{"_id": objId},
												   bson.M{"$set": obj})
	if err2 != nil {
		log.Println(err2)
		return nil, err
	}
	log.Printf("updated document in %s, id: %s", coll, id)
	return []byte("success"), nil
}

func (self *Repository) Delete(coll string, deleteMode DeleteMode, id string) ([]byte, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalf("Delete %s", err)
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
		return nil, err
	}

	return []byte("success"), nil
}