package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/bfbarry/CollabSource/back-end/errors"
	"github.com/bfbarry/CollabSource/back-end/repository"
)

const projectCollection = "projects"

type ProjectHandler struct {
	repository *repository.Repository
}

func BuildProjectHandler() *ProjectHandler { //TODO: init
	log.Println("Building project handler")
	projectHandler := &ProjectHandler{repository: repository.GetMongoRepository()}
	return projectHandler
}

func (self *ProjectHandler) CreateProject(streamObj *io.ReadCloser) (*Resp, *errors.Error) {
	msg, err := self.repository.Insert(projectCollection, streamObj)
	if err != nil {	
		return nil, err
	}
	// TODO: get id and return obj
	return newResponse(http.StatusCreated, msg), nil
}

func (self *ProjectHandler) GetProjectByID(id string) (*Resp, *errors.Error) {
	var op errors.Op = "controllers.GetProjectByID"

	result, err := self.repository.FindByID(projectCollection, id)
	if err != nil { // TODO: 400
		return nil, err
	}

	jsonResponse, jsonerr := json.Marshal(result)

	if jsonerr != nil { // TODO: handle json with function
		return nil, errors.E(err, http.StatusInternalServerError, op, "json marshall error")
	}

	return newResponse(http.StatusOK, jsonResponse), nil
}

func (self *ProjectHandler) GetProjectsByFilter(streamFilterObj *io.ReadCloser, pageNumber int64, pageSize int64) (*Resp, *errors.Error) {
	var op errors.Op = "controllers.GetProjectsByFilter"
	results, err := self.repository.Find(projectCollection, streamFilterObj, 0, 10)
	if err != nil {
		return nil, err
	}
	var jsonErr error
	jsonResponse, jsonErr := json.Marshal(results)
	if err != nil { // TODO: 500
		return nil, errors.E(jsonErr, http.StatusInternalServerError, op, "json marshall error")
	}
	return newResponse(http.StatusOK, jsonResponse), nil
}

func (self *ProjectHandler) UpdateProject(id string, streamObj *io.ReadCloser) (*Resp, *errors.Error) {
	_, err := self.repository.Update(projectCollection, streamObj, id)
	if err != nil {
		return nil, err
	}
	return newResponse(http.StatusOK,[]byte("success")), nil
}

func (self *ProjectHandler) DeleteProject(deleteMode repository.DeleteMode, id string) (*Resp, *errors.Error) {
	msg, err := self.repository.Delete(projectCollection, deleteMode, id)
	if err != nil {
		return nil, err
	}
	return newResponse(http.StatusOK, msg), nil
}