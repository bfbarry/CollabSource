package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Model interface{}

type UserCheck struct {
	Email string `json:"email"       bson:"email,omitempty"`
}

type PublicUser struct {
	Name        string   `json:"name"        bson:"name,omitempty"`
	Description string   `json:"description" bson:"description,omitempty"`
	Skills      []string `json:"skills"      bson:"skills,omitempty"`
}

type User struct {
	Id          primitive.ObjectID `json:"_id"        bson:"_id,omitempty"`
	Name        string             `json:"name"        bson:"name,omitempty"`
	Email       string             `json:"email"       bson:"email,omitempty"`
	Password    string             `json:"password"    bson:"password,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Skills      []string           `json:"skills"      bson:"skills,omitempty"`
}

type ProjectCheck struct {
	OwnerEmail string `json:"ownerEmail"  bson:"ownerEmail,omitempty"`
}

type Project struct {
	Id          primitive.ObjectID   `json:"id"        bson:"_id,omitempty"`
	Name        string   `json:"name"        bson:"name,omitempty"`
	Description string   `json:"description" bson:"description,omitempty"`
	Category    string   `json:"category"    bson:"category,omitempty"`
	Tags        []string `json:"tags"        bson:"tags,omitempty"`
	OwnerEmail  string   `json:"ownerEmail"  bson:"ownerEmail,omitempty"`
	// DateCreated string
	// Creator string
	// Admins []string
	//  []string
	// Location    string   `json:"location"`
}
