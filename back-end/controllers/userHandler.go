package controllers

import (
	"github.com/bfbarry/CollabSource/back-end/repository"
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

func (self *UserHandler) GetAllUsers(){
	self.repository.FindAll(userCollection)
}

func (self *UserHandler) GetUser(){
	self.repository.FindOne(userCollection)
}

func (self *UserHandler) CreateUser(){
	self.repository.Insert(userCollection)
}
