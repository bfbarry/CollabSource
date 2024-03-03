package model

type Model interface {}

type User struct {
	Name        string   `json:"name"`   
	Email       string   `json:"email"`     
}

type Project struct {
	Name        string   `json:"name"        bson:"name,omitempty"`
	Description string   `json:"description" bson:"description,omitempty"`
	Category 	string   `json:"category"    bson:"category,omitempty"`
	Tags        []string `json:"tags"        bson:"tags,omitempty"`
	// DateCreated string
	// Members []string
	// Location    string   `json:"location"`
}

func NewProject() *Project {
	return &Project{}
}

func NewUser() *User {
	return &User{}
}

func GetModelFromName(name string) Model {
	switch name {
		case "project":
			return NewProject()
		case "user":
			return NewUser()
		default:
			return nil
	}
}