package controllers

import (
	"github.com/bfbarry/CollabSource/back-end/repository"
	// "github.com/bfbarry/CollabSource/back-end/model"
	"fmt"
)

const userCollection = "users"

type UserHandler struct{
	repository repository.Repository
}

func BuildUserHandler() UserHandler{
	fmt.Println("Building user handler")
	userHandler := UserHandler{repository: *repository.GetMongoRepository()}
	return userHandler
}

// func (self *UserHandler) GetAllUsers(){
// 	self.repository.Find(userCollection)
// }

func (self *UserHandler) GetUser(id string){
	// self.repository.FindByID(userCollection ,id)
	return
}

// func (self *UserHandler) CreateUser(){
// 	u := model.User{} // TODO
// 	self.repository.Insert(userCollection, u)
// }
