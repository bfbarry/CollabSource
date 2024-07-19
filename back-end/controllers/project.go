package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"log"
	"strings"
	// "github.com/bfbarry/CollabSource/back-end/errors"
	"reflect"

	"github.com/bfbarry/CollabSource/back-end/model"
	"github.com/bfbarry/CollabSource/back-end/repository"
	"github.com/bfbarry/CollabSource/back-end/responseEntity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PROJECT_COLLECTION = "projects"

type ProjectController struct {
	repository *repository.Repository
}

var defaultProjectController *ProjectController

func GetProjectController() *ProjectController {
	return defaultProjectController
}

func init() {
	defaultProjectController = &ProjectController{repository: repository.GetMongoRepository()}
}

func (self *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {

	projectEntity := model.Project{}
	if err := json.NewDecoder(r.Body).Decode(&projectEntity); err != nil {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("Invalid JSON"))
		return
	}

	if projectEntity.Name == "" || projectEntity.Description == "" {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid payload"))
		return
	}

	if err := self.repository.Insert(PROJECT_COLLECTION, projectEntity); err != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("server error on insert"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, []byte("Success"))
}

func (self *ProjectController) GetProjectByID(w http.ResponseWriter, id string) {
	// var op errors.Op = "controllers.GetProjectByID"
	projectEntity := &model.Project{}

	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid Object ID"))
		return
	}

	mongoErr := self.repository.FindByID(PROJECT_COLLECTION, ObjId, projectEntity)
	//error handling for StatusNotFound
	if reflect.DeepEqual(*projectEntity, model.Project{}) {
		responseEntity.SendRequest(w, http.StatusNotFound, []byte("not found"))
		return
	} else if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	jsonResponse, jsonerr := json.Marshal(projectEntity)
	if jsonerr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, jsonResponse)
}

func (self *ProjectController) UpdateProject(w http.ResponseWriter, id string, r *http.Request) {

	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid Object ID"))
		return
	}

	projectEntity := model.Project{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&projectEntity); err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "json: unknown field") {
			responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Unexpected fields in JSON"))
			return
		}
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("something went wrong"))
		return
	}

	if reflect.DeepEqual(projectEntity, model.Project{}) {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("Invalid empty JSON"))
		return
	}

	exists, err := self.repository.DocumentExists(PROJECT_COLLECTION, ObjId) 
	if err != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}
	if !exists {
		responseEntity.SendRequest(w, http.StatusNotFound, []byte("project not found"))
		return
	}
	updatedCount, mongoErr := self.repository.Update(PROJECT_COLLECTION, ObjId, projectEntity)
	if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	if updatedCount == 0 {
		responseEntity.SendRequest(w, http.StatusNoContent, []byte("no change"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, []byte("success"))
}

func (self *ProjectController) DeleteProject(w http.ResponseWriter, id string) {
	// TODO pass in reader to get URL param

	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid Object ID"))
		return
	}
	exists, err := self.repository.DocumentExists(PROJECT_COLLECTION, ObjId) 
	if err != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}
	if !exists {
		responseEntity.SendRequest(w, http.StatusNotFound, []byte("project not found"))
		return
	}
	// deleteModeStr := r.URL.Query().Get("mode") // TODO separate hard and soft delete in repository.go
	// deleteMode := repository.Str2Enum(deleteModeStr)
	_, mongoErr := self.repository.Delete(PROJECT_COLLECTION, ObjId)
	if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}


	responseEntity.SendRequest(w, http.StatusOK, []byte("Success"))
}

func (self *ProjectController) GetProjects(w http.ResponseWriter, r *http.Request) {
	defaultPageNum := 1
	defaultPageSize := 10
	var pageNum int
	var pageSize int
	var err error

	queryParams := r.URL.Query()

	if pageNum, err = strconv.Atoi(queryParams.Get("page")); err != nil {
		pageNum = defaultPageNum
	}

	if pageSize, err = strconv.Atoi(queryParams.Get("size")); err != nil {
		pageSize = defaultPageSize
	}

	var projectEntity []model.Project
	filt := bson.M{}
	mongoErr := self.repository.FindManyByPage(PROJECT_COLLECTION, &projectEntity, pageNum, pageSize, filt)
	if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	response := responseEntity.PaginatedResponseBody[model.Project]{
		Data: projectEntity,
		Page: pageNum,
	}

	jsonResponse, jsonerr := json.Marshal(response)
	if jsonerr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, jsonResponse)
}
