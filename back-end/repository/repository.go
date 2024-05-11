package repository

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/bfbarry/CollabSource/back-end/errors"
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

func (self *Repository) Insert(coll string, streamObj *io.ReadCloser) ([]byte, *errors.Error) {
	var op errors.Op = "repository.Insert"
	obj, err := streamToObj(coll, streamObj, true)
	if err != nil {
		return nil, err
	}
	res, mongoerr := self.getCollection(coll).InsertOne(context.TODO(), obj)
	if mongoerr != nil {
		// TODO: once we know error space, switch and return them
		// var writeErr mongo.WriteException
		// if errors.As(err, &writeErr) {
		// 	log.Println("WriteException in Insert")
		// 	for _, we := range writeErr.WriteErrors {
		// 		log.Println(we)
		// 	}
		// } else {
		// }
		return nil, errors.E(mongoerr, http.StatusBadRequest, op, "bad insert")
	}
	log.Printf("inserted document with ID %v\n", res.InsertedID)
	return []byte("success"), nil
}

func (self *Repository) FindByID(coll string, id string) (model.Model, *errors.Error) {
	var op errors.Op = "repository.FindByID"
	obj := model.GetModelFromName(coll)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.E(err, http.StatusBadRequest, op, "not a valid ObjectID")
	}
	filter := bson.M{"_id": objId}
	err = self.getCollection(coll).FindOne(context.TODO(), filter).Decode(obj)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, errors.E(err, http.StatusBadRequest, op, "no documents found")
		default: // TODO define larger error space
			return nil, errors.E(err, http.StatusInternalServerError, op, "")
		}
	}
	return obj, nil
}

//filter e.g, bson.M{"category": filterField}
//		bson.M{} for no filter
// TODO: set limit on pageSize
func (self *Repository) Find(coll string, streamFilterObj *io.ReadCloser, pageIndex int64, pageSize int64) ([]model.Model, *errors.Error){
	var op errors.Op = "repository.Find"

	filter, err := streamToBsonM(coll, streamFilterObj)
	if err != nil {
		return nil, err
	}
	findOptions := options.Find()
	skip := pageIndex * pageSize
	findOptions.SetSkip(skip)
	findOptions.SetLimit(pageSize)
	cursor, findErr := self.getCollection(coll).Find(context.TODO(), filter, findOptions)
	if findErr != nil {
		log.Println(findErr)
		return nil, errors.E(findErr, http.StatusBadRequest, op, "no documents found")
	}

	results, sliceErr := cursorToSlice(cursor, coll)
	if sliceErr != nil {
		log.Println(sliceErr)
		return nil, errors.E(sliceErr, http.StatusInternalServerError, op, "")
	}
	return results, nil
}

func (self *Repository) Update(coll string, streamObj *io.ReadCloser, id string) ([]byte, *errors.Error) {
	var op errors.Op = "repository.Update"
	obj, err := streamToObj(coll, streamObj, false)
	if err != nil {
		return nil, err
	}
	objId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		return nil, errors.E(errId, http.StatusBadRequest, op, "not a valid ObjectID")
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

func (self *Repository) Delete(coll string, deleteMode DeleteMode, id string) ([]byte, *errors.Error) {
	var op errors.Op = "repository.Delete"
	objId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		return nil, errors.E(errId, http.StatusBadRequest, op, "not a valid ObjectID")
	}
	var del_err error
	switch deleteMode {
		case SoftDelete:
			_, del_err = self.getCollection(coll).UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": bson.M{"deleted": true}})
		case HardDelete:
			_, del_err = self.getCollection(coll).DeleteOne(context.TODO(), bson.M{"_id": objId})
		}
	if del_err != nil {	
		return nil, errors.E(del_err, http.StatusBadRequest, op, "Could not delete document")
	}

	return []byte("success"), nil
}