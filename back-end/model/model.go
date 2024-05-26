package model

type Model interface{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Project struct {
	// Id        string   `json:"_id"        bson:"_id,omitempty"`
	Name        string   `json:"name"        bson:"name,omitempty"`
	Description string   `json:"description" bson:"description,omitempty"`
	Category    string   `json:"category"    bson:"category,omitempty"`
	Tags        []string `json:"tags"        bson:"tags,omitempty"`
	// DateCreated string
	// Members []string
	// Location    string   `json:"location"`
}
