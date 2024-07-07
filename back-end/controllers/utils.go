package controllers
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/bfbarry/CollabSource/back-end/model"
	"log"

)
//TODO build filter dynamically 
func postQueryToBsonM(query model.UserPostQuery) (bson.M, error) {
	strIds := query.IDs
	var objIDs []primitive.ObjectID
	objIDs, err := stringsToObjectIDs(strIds)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// TODO object wrapper to inject more fields depending on query
	filter := bson.M{"_id": bson.M{"$in": objIDs}}
	return filter, nil
}

func stringsToObjectIDs(strSlice []string ) ([]primitive.ObjectID, error) {
	var objectIDs []primitive.ObjectID

	for _, str := range strSlice {
		objID, err := primitive.ObjectIDFromHex(str)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, objID)
	}
	return objectIDs, nil
}