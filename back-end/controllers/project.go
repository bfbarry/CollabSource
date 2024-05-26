package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	// "github.com/bfbarry/CollabSource/back-end/errors"
	"github.com/bfbarry/CollabSource/back-end/model"
	"github.com/bfbarry/CollabSource/back-end/repository"
	"github.com/bfbarry/CollabSource/back-end/responseEntity"
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

	streamObj := r.Body
	projectEntity := model.Project{}
	if err := json.NewDecoder(streamObj).Decode(&projectEntity); err != nil {
		responseEntity.ResponseEntity(w, http.StatusBadRequest, []byte("Invalid JSON"))
		return
	}

	if projectEntity.Name == "" || projectEntity.Description == "" {
		responseEntity.ResponseEntity(w, http.StatusUnprocessableEntity, []byte("Invalid payload"))
		return
	}

	if err := self.repository.Insert(PROJECT_COLLECTION, projectEntity); err != nil {
		responseEntity.ResponseEntity(w, http.StatusInternalServerError, []byte("Error"))
		return
	}

	responseEntity.ResponseEntity(w, http.StatusOK, []byte("Success"))
}

func (self *ProjectController) GetProjectByID(w http.ResponseWriter, id string) {
	// var op errors.Op = "controllers.GetProjectByID"

	result, err := self.repository.FindByID(PROJECT_COLLECTION, id)
	if err != nil { // TODO: 400
		// return nil, err
	}

	jsonResponse, jsonerr := json.Marshal(result)

	if jsonerr != nil { // TODO: handle json with function
		// return nil, errors.E(err, http.StatusInternalServerError, op, "json marshall error")
	}

	responseEntity.ResponseEntity(w, http.StatusOK, jsonResponse)
	// return jsonResponse, nil
}

func (self *ProjectController) GetProjectsByFilter(w http.ResponseWriter, streamFilterObj *io.ReadCloser, pageNumber int64, pageSize int64) {
	// var op errors.Op = "controllers.GetProjectsByFilter"
	results, err := self.repository.Find(PROJECT_COLLECTION, streamFilterObj, 0, 10)
	if err != nil {
		// return nil, err
	}
	// var jsonErr error
	jsonResponse, _ := json.Marshal(results)
	if err != nil { // TODO: 500
		// return nil, errors.E(jsonErr, http.StatusInternalServerError, op, "json marshall error")
	}
	responseEntity.ResponseEntity(w, http.StatusOK, jsonResponse)
	// return jsonResponse, nil
}

func (self *ProjectController) UpdateProject(w http.ResponseWriter, id string, streamObj *io.ReadCloser) {
	_, err := self.repository.Update(PROJECT_COLLECTION, streamObj, id)
	if err != nil {
		// return nil, err
	}
	// return  []byte("success"), nil
	responseEntity.ResponseEntity(w, http.StatusOK, []byte("success"))
}

func (self *ProjectController) DeleteProject(w http.ResponseWriter, deleteMode repository.DeleteMode, id string) {
	msg, err := self.repository.Delete(PROJECT_COLLECTION, deleteMode, id)
	if err != nil {
		// return nil, err
	}
	// return msg, nil
	responseEntity.ResponseEntity(w, http.StatusOK, msg)
}
