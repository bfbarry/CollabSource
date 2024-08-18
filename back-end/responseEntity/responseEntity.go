package responseEntity

import (
	"net/http"
	// "fmt"
	// "encoding/json"
	// "github.com/bfbarry/CollabSource/back-end/errors"
	// "github.com/bfbarry/CollabSource/back-end/controllers"
)

func SendRequest(w http.ResponseWriter, status int, Body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(Body)
}

type PaginatedResponseBody[T any] struct {
	Items []T `json:"items"`
	Page int `json:"page"`
	HasNext bool `json:"hasNext"`
}
