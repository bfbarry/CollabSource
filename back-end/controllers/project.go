package controllers

import (
	"encoding/json"
	"io"
	"log"

	"github.com/bfbarry/CollabSource/back-end/repository"
)

const projectCollection = "projects"

type ProjectHandler struct {
	repository *repository.Repository
}

func BuildProjectHandler() *ProjectHandler {
	log.Println("Building project handler")
	projectHandler := &ProjectHandler{repository: repository.GetMongoRepository()}
	return projectHandler
}

func (self *ProjectHandler) CreateProject(streamObj *io.ReadCloser) ([]byte, error) {
	msg, err := self.repository.Insert(projectCollection, streamObj)
	if err != nil {	
		return nil, err
	}
	// TODO: get id and return obj
	return msg, nil
}

func (self *ProjectHandler) GetProjectByID(id string) ([]byte, error) {
	result, err := self.repository.FindByID(projectCollection, id)
	if err != nil { // TODO: 400
		return nil, err
	}
	jsonResponse, err := json.Marshal(result)

	if err != nil { // TODO: 500
		return nil, err
	}
	return jsonResponse, nil
}

func (self *ProjectHandler) GetProjectsByFilter(streamFilterObj *io.ReadCloser, pageNumber int64, pageSize int64) ([]byte, error) {
	results, err := self.repository.Find(projectCollection, streamFilterObj, 0, 10)
	if err != nil {
		return nil, err
	}
	jsonResponse, err := json.Marshal(results)
	if err != nil { // TODO: 500
		return nil, err
	}
	return jsonResponse, nil
}

func (self *ProjectHandler) UpdateProject(id string, streamObj *io.ReadCloser) ([]byte, error) {
	var err error
	_, err = self.repository.Update(projectCollection, streamObj, id)
	if err != nil {
		return nil, err
	}
	return []byte("success"), nil
}

func (self *ProjectHandler) DeleteProject(deleteMode repository.DeleteMode, id string) ([]byte, error) {
	msg, err := self.repository.Delete(projectCollection, deleteMode, id)
	if err != nil {
		return nil, err
	}
	return msg, nil
}