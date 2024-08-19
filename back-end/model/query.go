package model

// structs for json query
// Add fields as necessary
type LoginFields struct {
	Email string
	Password string
}

type UserPostQuery struct {
	IDs []string
}

type ProjectFilter struct {
	SearchQuery string     `json:"searchQuery" bson:"searchQuery,omitempty"`
	Categories  []string   `json:"categories" bson:"categories,omitempty"`
	// Tags        []string   `json:"tags" bson:"tags,omitempty"`
	// Seeking     []string 			 `json:"seeking" bson:"seeking,omitempty"`
}