package repository

import (
	"context"
	"io"
	"log"

	"github.com/bfbarry/CollabSource/back-end/errors"
	"github.com/bfbarry/CollabSource/back-end/model"
	"github.com/bfbarry/CollabSource/back-end/mongoClient"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct {
	mongoClient *mongo.Database
}

var mongoRepository *Repository

func GetMongoRepository() *Repository {
	return mongoRepository
}

func init() {
	// defer mongoClient.CloseMongoClient()
	mongoRepository = &Repository{mongoClient: mongoClient.GetMongoDb()}
}

func (self *Repository) getCollection(coll string) *mongo.Collection {
	return self.mongoClient.Collection(coll)
}

func (self *Repository) Insert(coll string, obj model.Model) *errors.Error {

	res, mongoerr := self.getCollection(coll).InsertOne(context.TODO(), obj)
	if mongoerr != nil {
		log.Printf("Error inserting object e message: %s", mongoerr)
		return &errors.Error{}
	}
	log.Printf("inserted document with ID %v\n", res.InsertedID)
	return nil
}

func (self *Repository) FindByID(coll string, id primitive.ObjectID, obj interface{}) (model.Model, *errors.Error) {
	// var op errors.Op = "repository.FindByID"
	
	filter := bson.M{"_id": id}
	err := self.getCollection(coll).FindOne(context.TODO(), filter).Decode(obj)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, nil
		default:
			return nil, &errors.Error{}
		}
	}
	return obj, nil
}

func (self *Repository) Update(coll string, id primitive.ObjectID, obj model.Model) (int64, *errors.Error) {
	// var op errors.Op = "repository.Update"
	
	result, err := self.getCollection(coll).UpdateOne(context.TODO(),bson.M{"_id": id},bson.M{"$set": obj})
	if err != nil {
		return 0, &errors.Error{}
	}
	
	return result.ModifiedCount, nil
}

func (self *Repository) Delete(coll string, id primitive.ObjectID) (int64, *errors.Error) {
	//var op errors.Op = "repository.Delete"
	
	// var del_err error
	// switch deleteMode {
	// case SoftDelete:
	// 	_, del_err = self.getCollection(coll).UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": bson.M{"deleted": true}})
	// case HardDelete:
	result, err := self.getCollection(coll).DeleteOne(context.TODO(), bson.M{"_id": id})
	// }
	if err != nil {
		return 0, &errors.Error{}
	}

	return result.DeletedCount, nil
}

// filter e.g, bson.M{"category": filterField}
//
//	bson.M{} for no filter
//
// TODO: set limit on pageSize
func (self *Repository) Find(coll string, streamFilterObj *io.ReadCloser, pageIndex int64, pageSize int64) ([]model.Model, *errors.Error) {
	// var op errors.Op = "repository.Find"

	// filter, err := streamToBsonM(coll, streamFilterObj)
	// if err != nil {
	// 	return nil, err
	// }
	// findOptions := options.Find()
	// skip := pageIndex * pageSize
	// findOptions.SetSkip(skip)
	// findOptions.SetLimit(pageSize)
	// cursor, findErr := self.getCollection(coll).Find(context.TODO(), filter, findOptions)
	// if findErr != nil {
	// 	log.Println(findErr)
	// 	return nil, errors.E(findErr, http.StatusBadRequest, op, "no documents found")
	// }

	// results, sliceErr := cursorToSlice(cursor, coll)
	// if sliceErr != nil {
	// 	log.Println(sliceErr)
	// 	return nil, errors.E(sliceErr, http.StatusInternalServerError, op, "")
	// }
	// return results, nil
	return nil, nil
}
