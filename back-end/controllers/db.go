package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)
// thanks https://www.alexedwards.net/blog/organising-database-access
// used for dependency injection
type Env struct {
	DB *mongo.Database
}