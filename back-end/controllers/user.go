package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct{

}

type UserEnv struct {
	Coll  *mongo.Collection
}