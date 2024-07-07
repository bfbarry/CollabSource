package controllers

import (
	"encoding/json"
	"strconv"
	"net/http"

	"github.com/bfbarry/CollabSource/back-end/model"
	"github.com/bfbarry/CollabSource/back-end/repository"
	"github.com/bfbarry/CollabSource/back-end/responseEntity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

const USER_COLLECTION = "users"

type UserController struct {
	repository *repository.Repository
}

var defaultUserController *UserController

func GetUserController() *UserController {
	return defaultUserController
}

func init() {
	defaultUserController = &UserController{repository: repository.GetMongoRepository()}
}

func (self *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	streamObj := r.Body
	userEntity := model.User{}

	if err := json.NewDecoder(streamObj).Decode(&userEntity); err != nil {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("Invalid JSON"))
		return
	}

	if userEntity.Name == "" || userEntity.Description == "" {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid payload"))
		return
	}

	if err := self.repository.Insert(USER_COLLECTION, userEntity); err != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("server error on insert"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, []byte("success"))
}

func (self *UserController) GetUserByID(w http.ResponseWriter, id string) {
	userEntity := &model.User{}
	ObjId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid Object ID"))
		return
	}

	result, mongoErr := self.repository.FindByID(USER_COLLECTION, ObjId, userEntity)
	if mongoErr != nil {
		//TODO log
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}
	if result == nil {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("ID Not Found"))
		return
	}

	jsonRes, jsonerr := json.Marshal(result)
	if jsonerr != nil { 
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, jsonRes)
}

func (self *UserController) GetUsersByQuery(w http.ResponseWriter, r *http.Request) {
	defaultPageNum := 1
	defaultPageSize := 10
	var err error
	query := model.UserPostQuery{}
	
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("Invalid JSON"))
		return
	}
	log.Println(query)
	var filter bson.M
	filter, err = postQueryToBsonM(query)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("Invalid JSON"))
		return
	}

	var pageNum int
	var pageSize int

	queryParams := r.URL.Query()

	if pageNum, err = strconv.Atoi(queryParams.Get("page")); err != nil {
		pageNum = defaultPageNum
	}

	if pageSize, err = strconv.Atoi(queryParams.Get("size")); err != nil {
		pageNum = defaultPageSize
	}

	var userEntities []model.User

	mongoErr := self.repository.FindManyByPage(USER_COLLECTION, &userEntities, pageNum, pageSize, filter)
	if mongoErr != nil {
		log.Println(mongoErr)
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}
	res := responseEntity.PaginatedResponseBody[model.User] {
		Data: userEntities,
		Page: pageNum,
	}

	jsonRes, jsonErr := json.Marshal(res)
	if jsonErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, jsonRes)
}

func (self *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	defaultPageNum := 1
	defaultPageSize := 10
	var err error

	var pageNum int
	var pageSize int

	queryParams := r.URL.Query()

	if pageNum, err = strconv.Atoi(queryParams.Get("page")); err != nil {
		pageNum = defaultPageNum
	}

	if pageSize, err = strconv.Atoi(queryParams.Get("size")); err != nil {
		pageNum = defaultPageSize
	}

	var userEntities []model.User
	filter := bson.M{}

	mongoErr := self.repository.FindManyByPage(USER_COLLECTION, &userEntities, pageNum, pageSize, filter)
	if mongoErr != nil {
		log.Println(mongoErr)
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}
	res := responseEntity.PaginatedResponseBody[model.User] {
		Data: userEntities,
		Page: pageNum,
	}

	jsonRes, jsonErr := json.Marshal(res)
	if jsonErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, jsonRes)
}

func (self *UserController) UpdateUser(w http.ResponseWriter, id string, r *http.Request) {

	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid Object ID"))
		return
	}

	streamObj := r.Body
	userEntity := model.User{}
	if err := json.NewDecoder(streamObj).Decode(&userEntity); err != nil {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("Invalid JSON"))
		return
	}

	updatedCount, mongoErr := self.repository.Update(USER_COLLECTION, ObjId, userEntity)
	if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	if updatedCount == 0 {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("ID Not Found"))
		return
	}
	
	responseEntity.SendRequest(w, http.StatusOK, []byte("success"))
}

func (self *UserController) DeleteUser(w http.ResponseWriter, id string) {
	// TODO pass in reader to get URL param

	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid Object ID"))
		return
	}

	// deleteModeStr := r.URL.Query().Get("mode") // TODO separate hard and soft delete in repository.go
	// deleteMode := repository.Str2Enum(deleteModeStr)
	deletedCount, mongoErr := self.repository.Delete(USER_COLLECTION, ObjId)
	if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	if deletedCount == 0{
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("ID Not Found"))
		return
	}
	
	responseEntity.SendRequest(w, http.StatusOK, []byte("Success"))
}