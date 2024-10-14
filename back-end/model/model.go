package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*  NOTE
struct tags should be separated by one space.
*/

type Model interface{}

type IdObj struct {
	Id primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
}

type PublicUser struct {
	Name        string   `json:"name" bson:"name,omitempty"`
	Description string   `json:"description" bson:"description,omitempty"`
	Skills      []string `json:"skills" bson:"skills,omitempty"`
}

type User struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Email       string             `json:"email" bson:"email,omitempty"`
	Password    string             `json:"password" bson:"password,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Skills      []string           `json:"skills" bson:"skills,omitempty"`
}

type ProjectCheck struct {
	OwnerId primitive.ObjectID `json:"ownerId" bson:"ownerId,omitempty"`
}

type Project struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	OwnerId        primitive.ObjectID `json:"ownerId" bson:"ownerId,omitempty"`
	OwnerEmail     string             `json:"ownerEmail" bson:"ownerEmail,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	Description    string             `json:"description" bson:"description,omitempty"`
	Category       string             `json:"category" bson:"category,omitempty"`
	Tags           []string           `json:"tags" bson:"tags,omitempty"`
	Seeking        []string           `json:"seeking" bson:"seeking,omitempty"`
	Members        []Member           `json:"members" bson:"members,omitempty"`
	DateCreated    time.Time          `json:"dateCreated" bson:"dateCreated,omitempty"`
	Links          []string           `json:"links" bson:"links,omitempty"`
	MemberRequests []MemberRequest    `json:"memberRequests" bson:"memberRequests,omitempty"`
	// Admins []string
	// Location    string   `json:"location"`
}

type ProjectResponse struct {
	Project   Project `json:"project"`
	IsPublic  bool    `json:"isPublic"`
	IsMember  bool    `json:"isMember"`
	IsOwner   bool    `json:"isOwner"`
	IsPending bool    `json:"isPending"`
}

func (p *Project) BuildProjectResponse(isPublic bool, isMember bool, isOwner bool, isPending bool) ProjectResponse {
	var project = Project{}
	project.Id = p.Id
	project.OwnerId = p.OwnerId
	project.Name = p.Name
	project.Description = p.Description
	project.Category = p.Category
	project.Tags = p.Tags
	project.Seeking = p.Seeking
	project.DateCreated = p.DateCreated
	if isMember || isOwner {
		project.Members = p.Members
		project.Links = p.Links
	}
	if isOwner {
		requestArray := []MemberRequest{}
		for _, value := range p.MemberRequests {
			value.Email = ""
			requestArray = append(requestArray, value)
		}
		project.MemberRequests = requestArray
	}

	return ProjectResponse{
		Project:   project,
		IsPublic:  isPublic,
		IsMember:  isMember,
		IsOwner:   isOwner,
		IsPending: isPending,
	}
}

type PaginatedResponseBody[T any] struct {
	Data []T `json:"data"`
	Page int `json:"page"`
}

type ProjectRequest struct {
	// Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	ProjectId primitive.ObjectID `json:"projectId" bson:"projectId,omitempty"`
	// Message   string             `json:"message" bson:"message,omitempty"`
}

type Member struct {
	UserId string `json:"userId" bson:"userId,omitempty"`
	Name   string `json:"name" bson:"name,omitempty"`
	Email  string `json:"email" bson:"email,omitempty"`
}

type MemberRequest struct {
	UserId string `json:"userId" bson:"userId,omitempty"`
	Name   string `json:"name" bson:"name,omitempty"`
	Email  string `json:"email" bson:"email,omitempty"`
}

type MemberRequestResponse struct {
	UserId string `json:"userId" bson:"userId,omitempty"`
	Name   string `json:"name" bson:"name,omitempty"`
}

type ProjectRequestAdmission struct {
	Member
	Admission string `json:"admission"`
}
