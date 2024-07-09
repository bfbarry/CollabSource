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