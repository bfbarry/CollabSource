package tests

import (
	"testing"
	"net/http"

)
var testBaseUrl string

func init() {
	testBaseUrl = "http://127.0.0.1:8080/api/v1/project"
}

func TestGetProjectByID(t *testing.T) {
	var _ error
	validId := "65ad98995c7412fb601b8e44" // TODO: could get some random id automatically
	testGetRequest(testBaseUrl, validId, http.StatusOK)

	nonexistentId := "75ad98995c7412fb601b8e45" 
	testGetRequest(testBaseUrl, nonexistentId, http.StatusBadRequest)

	notObjectId := "aaaa" 
	testGetRequest(testBaseUrl, notObjectId, http.StatusBadRequest)

}

func TestCreateProject(t *testing.T) {
	goodBody := []byte(`{
		"name"        : "test",
		"description" : "test",
		"category" 	  : "test",
		"tags"        : ["test", "test"]
	}`)
	postRequest(testBaseUrl, "create", http.StatusCreated, goodBody)
	
	missingFieldBody := []byte(`{
		"description" : "test",
		"category" 	  : "test",
		"tags"        : ["test", "test"]
	}`)
	postRequest(testBaseUrl, "create", http.StatusBadRequest, missingFieldBody)

	badFieldBody := []byte(`{
		"name1"        : "test",
		"description" : "test",
		"category" 	  : "test",
		"tags"        : ["test", "test"]
	}`)
	postRequest(testBaseUrl, "create", http.StatusBadRequest, badFieldBody)
}