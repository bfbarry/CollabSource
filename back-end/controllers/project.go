package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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

func (self *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request, UUID string) {

	projectEntity := model.Project{}
	if err := json.NewDecoder(r.Body).Decode(&projectEntity); err != nil {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("Invalid JSON"))
		return
	}

	if projectEntity.Name == "" || projectEntity.Description == "" {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid payload"))
		return
	}
	projectEntity.DateCreated = time.Now()
	projectEntity.OwnerEmail = UUID
	id, err := self.repository.Insert(PROJECT_COLLECTION, projectEntity)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("server error on insert"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, []byte(id.Hex()))
}

func (self *ProjectController) GetProjectByID(w http.ResponseWriter, id string, userUUID string, userId string) {
	// var op errors.Op = "controllers.GetProjectByID"
	// TODO return different data if member/admin

	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid Object ID"))
		return
	}

	projectEntity := &model.Project{}

	mongoErr := self.repository.FindByID(PROJECT_COLLECTION, ObjId, projectEntity)
	//error handling for StatusNotFound
	if reflect.DeepEqual(*projectEntity, model.Project{}) {
		responseEntity.SendRequest(w, http.StatusNotFound, []byte("not found"))
		return
	} else if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	var jsonResponse []byte
	var jsonerr error

	if userUUID == "public" {
		jsonResponse, jsonerr = json.Marshal(projectEntity.BuildProjectResponse(true, false, false, false))
	} else if projectEntity.OwnerEmail == userUUID {
		jsonResponse, jsonerr = json.Marshal(projectEntity.BuildProjectResponse(false, false, true, false))
	} else {
		isMember := false
		for _, member := range projectEntity.Members {
			if member.UserId == userId {
				jsonResponse, jsonerr = json.Marshal(projectEntity.BuildProjectResponse(false, true, false, false))
				isMember = true
				break
			}
		}
		if !isMember {
			isPendingMember := false
			for _, pendingMember := range projectEntity.MemberRequests {
				if pendingMember.Email == userUUID {
					jsonResponse, jsonerr = json.Marshal(projectEntity.BuildProjectResponse(false, false, false, true))
					isPendingMember = true
					break
				}
			}
			if !isPendingMember {
				jsonResponse, jsonerr = json.Marshal(projectEntity.BuildProjectResponse(true, false, false, false))
			}
		}
	}

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
	// TODO use elasticsearch, add more fields to ProjectFilter
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
	
	projectFilter := model.ProjectFilter{}
	if err := json.NewDecoder(r.Body).Decode(&projectFilter); err != nil {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("Invalid JSON"))
		return
	}
	filt := bson.M{}

	if len(projectFilter.Categories) > 0 {
		filt["category"] = bson.M{"$in": projectFilter.Categories}
	}
	if len(projectFilter.SearchQuery) > 0 {
		splitTerms := strings.Fields(projectFilter.SearchQuery)
		tagConditions := make([]bson.M, len(splitTerms))
		for i, term := range(splitTerms) {
			tagConditions[i] = bson.M{
				"tags": bson.M{
					"$elemMatch": bson.M{
						"$regex": term,
						"$options": "i",
					},
				},
			}
		}
		tagFilt := bson.M{"$or": tagConditions}
		// tagFilt := bson.M{"tags": bson.M{"$in": splitTerms}}

		search := bson.M{"$or": []bson.M{
			{"description": bson.M{"$regex": projectFilter.SearchQuery,
									"$options": "i"}},
			{"name": bson.M{"$regex": projectFilter.SearchQuery,
									"$options": "i"}},
			tagFilt,
			},
		}

		//join any previous filters into one filter 
		if len(projectFilter.Categories) > 0 {
			filt = bson.M{"$and": []bson.M{search, filt}}
		} else {
			filt = search
		}
	}

	hasNext, mongoErr := self.repository.FindManyByPage(PROJECT_COLLECTION, &projectEntity, pageNum, pageSize, filt)
	if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong here"))
		return
	}

	if err != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}
	
	response := responseEntity.PaginatedResponseBody[model.Project]{
		Items: projectEntity,
		Page: pageNum,
		HasNext: hasNext,
	}

	jsonResponse, jsonerr := json.Marshal(response)
	if jsonerr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	responseEntity.SendRequest(w, http.StatusOK, jsonResponse)
}

func (self *ProjectController) SendProjectRequest(w http.ResponseWriter, r *http.Request, userUUID string) {
	if userUUID == "public" {
		responseEntity.SendRequest(w, http.StatusUnauthorized, []byte("unauthorized"))
		return
	}
	projectRequestEntity := model.ProjectRequest{}
	if err := json.NewDecoder(r.Body).Decode(&projectRequestEntity); err != nil {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("Invalid JSON"))
		return
	}

	userEntity := &model.User{}
	mongoErr := self.repository.FindByID(USER_COLLECTION, projectRequestEntity.UserId, userEntity)
	if reflect.DeepEqual(*userEntity, model.User{}) || userEntity.Email != userUUID {
		responseEntity.SendRequest(w, http.StatusNotFound, []byte("not found"))
		return
	}
	if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}

	request := bson.D{
		{Key: "userId", Value: projectRequestEntity.UserId.Hex()},
		{Key: "name", Value: userEntity.Name},
		{Key: "email", Value: userUUID},
	}

	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "memberRequests", Value: request}},
		}}
	updatedCount, updateMongoErr := self.repository.UpdateWithKeys(PROJECT_COLLECTION, projectRequestEntity.ProjectId, update)
	if updateMongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}
	if updatedCount == 0 {
		responseEntity.SendRequest(w, http.StatusNoContent, []byte("no change"))
		return
	}

	responseEntity.SendRequest(w, http.StatusCreated, []byte("Success"))
}

func (self *ProjectController) RespondToProjectRequest(w http.ResponseWriter, r *http.Request, id string, userUUID string) {

	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseEntity.SendRequest(w, http.StatusUnprocessableEntity, []byte("Invalid Object ID"))
		return
	}

	project := &model.Project{}
	var mongoErr error
	mongoErr = self.repository.FindByID(PROJECT_COLLECTION, ObjId, project)
	if reflect.DeepEqual(*project, model.Project{}) {
		responseEntity.SendRequest(w, http.StatusNotFound, []byte("project not found"))
		return
	}
	if mongoErr != nil {
		responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
		return
	}
	if project.OwnerEmail != userUUID {
		responseEntity.SendRequest(w, http.StatusUnauthorized, []byte("unauthorized"))
		return
	}

	projectRequestAdmission := model.ProjectRequestAdmission{}
	if err := json.NewDecoder(r.Body).Decode(&projectRequestAdmission); err != nil {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("Invalid JSON"))
		return
	}

	if projectRequestAdmission.Admission == "accepted" {

		removeUpdate := bson.D{
			{Key: "$pull", Value: bson.D{
				{Key: "memberRequests", Value: bson.D{
					{Key: "userId", Value: projectRequestAdmission.UserId},
					{Key: "name", Value: projectRequestAdmission.Name},
				}},
			}},
		}

		var updatedCount int64
		var updateMongoErr error
		updatedCount, updateMongoErr = self.repository.UpdateWithKeys(PROJECT_COLLECTION, ObjId, removeUpdate)
		if updateMongoErr != nil {
			responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
			return
		}
		if updatedCount == 0 {
			responseEntity.SendRequest(w, http.StatusNoContent, []byte("no change"))
			return
		}

		updateRequest := bson.D{
			{Key: "userId", Value: projectRequestAdmission.UserId},
			{Key: "name", Value: projectRequestAdmission.Name},
		}

		insertUpdate := bson.D{
			{Key: "$push", Value: bson.D{
				{Key: "members", Value: updateRequest}},
			}}

		updatedCount, updateMongoErr = self.repository.UpdateWithKeys(PROJECT_COLLECTION, ObjId, insertUpdate)
		if updateMongoErr != nil {
			responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
			return
		}
		if updatedCount == 0 {
			responseEntity.SendRequest(w, http.StatusNoContent, []byte("no change"))
			return
		}

		responseEntity.SendRequest(w, http.StatusOK, []byte("Success"))

	} else if projectRequestAdmission.Admission == "denied" {

		update := bson.D{
			{Key: "$pull", Value: bson.D{
				{Key: "memberRequests", Value: bson.D{
					{Key: "userId", Value: projectRequestAdmission.UserId},
					{Key: "name", Value: projectRequestAdmission.Name},
				}},
			}},
		}

		updatedCount, updateMongoErr := self.repository.UpdateWithKeys(PROJECT_COLLECTION, ObjId, update)
		if updateMongoErr != nil {
			responseEntity.SendRequest(w, http.StatusInternalServerError, []byte("Something went wrong"))
			return
		}
		if updatedCount == 0 {

			responseEntity.SendRequest(w, http.StatusOK, []byte("no change"))
			return
		}
		responseEntity.SendRequest(w, http.StatusOK, []byte("Success"))

	} else {
		responseEntity.SendRequest(w, http.StatusBadRequest, []byte("Invalid Payload"))
	}
}

// func (self *ProjectController) GetProjectRequests(w http.ResponseWriter, r *http.Request, projectId string, userUUID string) {
// 	//check if UUID == Project(projectId).OwnerId
// }
