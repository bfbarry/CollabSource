package repository

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/bfbarry/CollabSource/back-end/errors"
	"github.com/bfbarry/CollabSource/back-end/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

func structValueIsEmpty(val reflect.Value) bool {
	//TODO: return error specific to type
	switch val.Kind() {
	case reflect.String, reflect.Array:
		return val.Len() == 0
	case reflect.Map, reflect.Slice:
        return val.Len() == 0 || val.IsNil()
    case reflect.Bool:
        return !val.Bool()
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return val.Int() == 0
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        return val.Uint() == 0
    case reflect.Float32, reflect.Float64:
        return val.Float() == 0
    case reflect.Interface, reflect.Ptr:
        return val.IsNil()
	}
	return true
}

func validateFullStruct(obj model.Model) bool {
	// var op errors.Op = "repository.validateFullStruct"
	// obj := model.GetModelFromName(coll)
	reflVal := reflect.ValueOf(obj).Elem()
	for i:=0; i < reflVal.NumField(); i++ {
		return !structValueIsEmpty(reflVal.Field(i))
	}
	return true

}

func streamToObj(coll string, reqBody *io.ReadCloser, insert bool) (model.Model, *errors.Error) {
	var op errors.Op = "repository.streamToObj"
	obj := model.GetModelFromName(coll)
	err := json.NewDecoder(*reqBody).Decode(&obj)
	//TODO: validate fields for update operation
	if insert && !validateFullStruct(obj) {
		return nil, errors.E(errors.New("Non full struct"), 
							http.StatusBadRequest, op, "Object must have fully populated fields for Insert")
	}
	if err != nil { //TODO: 400
		return nil, errors.E(err, http.StatusInternalServerError, op, "Bad decoding")
	}
	return obj, nil
}

func streamToBsonM(coll string, reqBody *io.ReadCloser) (bson.M, *errors.Error) {
	obj, err := streamToObj(coll, reqBody, false)
	if err != nil {
		return nil, err
	}
	bsonBytes, bsonerr := bson.Marshal(obj)
	if bsonerr != nil {
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