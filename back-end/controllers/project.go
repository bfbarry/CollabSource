package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	id, err := self.repository.Insert(PROJECT_COLLECTION, projectEntity)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("server error on insert"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, []byte(id.Hex()))
}

func (self *ProjectController) GetProjectByID(w http.ResponseWriter, id string, userUUID string) {
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

	if projectEntity.OwnerId != userUUID {
		projectEntity.OwnerId = ""
	}

	jsonResponse, jsonerr := json.Marshal(projectEntity)
	if jsonerr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, jsonResponse)
}

func (self *ProjectController) UpdateProject(w http.ResponseWriter, id string, r *http.Request, userUUID string) {

	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid Object ID"))
		return
	}

	ProjectCheck := &model.ProjectCheck{}
	var mongoErr error
	mongoErr = self.repository.FindByID(PROJECT_COLLECTION, ObjId, ProjectCheck)
	if reflect.DeepEqual(*ProjectCheck, model.ProjectCheck{}) {
		responseEntity.SendRequest(w, http.StatusNotFound, []byte("project not found"))
		return
	}
	if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}
	if ProjectCheck.OwnerId.String() != userUUID {
		responseEntity.SendRequest(w, http.StatusUnauthorized, []byte("unauthorized"))
		return
	}

	projectEntity := model.Project{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&projectEntity); err != nil {
		fmt.Println(err)
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

func (self *ProjectController) DeleteProject(w http.ResponseWriter, id string, userUUID string) {
	// TODO pass in reader to get URL param

	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid Object ID"))
		return
	}

	ProjectCheck := &model.ProjectCheck{}
	var mongoErr error
	mongoErr = self.repository.FindByID(PROJECT_COLLECTION, ObjId, ProjectCheck)
	if reflect.DeepEqual(*ProjectCheck, model.ProjectCheck{}) {
		responseEntity.SendRequest(w, http.StatusNotFound, []byte("project not found"))
		return
	}
	if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}
	if ProjectCheck.OwnerId.String() != userUUID {
		responseEntity.SendRequest(w, http.StatusUnauthorized, []byte("unauthorized"))
		return
	}

	// deleteModeStr := r.URL.Query().Get("mode") // TODO separate hard and soft delete in repository.go
	// deleteMode := repository.Str2Enum(deleteModeStr)
	_, mongoErr = self.repository.Delete(PROJECT_COLLECTION, ObjId)
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
