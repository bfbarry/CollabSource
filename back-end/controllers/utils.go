package controllers
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func stringSlice2ObjectIDSlice(strSlice []string ) ([]primitive.ObjectID, error) {
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