package tests

import (
	"fmt"
	"net/http"
	"testing"
)
var testBaseUrl string

func init() {
	testBaseUrl = "http://127.0.0.1:8080/api/v1/project"
}

func TestGetProjectByID(t *testing.T) {
	var _ error
	validId := "/65ad98995c7412fb601b8e44" // TODO: could get some random id automatically
	testGetRequest(testBaseUrl, validId, http.StatusOK)

	nonexistentId := "/75ad98995c7412fb601b8e45" 
	testGetRequest(testBaseUrl, nonexistentId, http.StatusBadRequest)

	notObjectId := "/aaaa" 
	testGetRequest(testBaseUrl, notObjectId, http.StatusBadRequest)

}

func TestCreateProject(t *testing.T) {
	goodBody := []byte(`{
		"name"        : "test",
		"description" : "test",
		"category" 	  : "test",
		"tags"        : ["test", "test"]
	}`)
	requestWithBody("POST", testBaseUrl, "/create", http.StatusCreated, goodBody)
	
	missingFieldBody := []byte(`{
		"description" : "test",
		"category" 	  : "test",
		"tags"        : ["test", "test"]
	}`)
	requestWithBody("POST", testBaseUrl, "/create", http.StatusBadRequest, missingFieldBody)

	badFieldBody := []byte(`{
		"name1"        : "test",
		"description" : "test",
		"category" 	  : "test",
		"tags"        : ["test", "test"]
	}`)
	requestWithBody("POST", testBaseUrl, "/create", http.StatusBadRequest, badFieldBody)
}

func TestUpdateProject(t *testing.T) {
	validId := "/65ad98995c7412fb601b8e44" // TODO: could get some random id automatically
	goodBody := []byte(`{
		"name"        : "UPDATED NAME",
		"category" 	  : "test",
		"tags"        : ["test", "test"]
	}`)
	requestWithBody("PATCH", testBaseUrl, validId, http.StatusOK, goodBody)
	
	// TODO: need to test failure
	
}

func TestDeleteProject(t *testing.T){
	validId := "65fb3d735513beede5b9bbf3"
	subUrl := fmt.Sprintf("/%s%s", validId, "?mode=soft")
	testDeleteRequest(testBaseUrl, subUrl, 200)
}

func TestGetProjects(t *testing.T) {
	goodBody := []byte(`{
		"category" 	  : "test"
	}`)
	requestWithBody("GET", "http://127.0.0.1:8080/api/v1/projects", "?pageNum=0&pageSize=10", http.StatusOK, goodBody)
}