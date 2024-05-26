package responseEntity

import (
	"net/http"
	// "log"
	// "encoding/json"
	// "github.com/bfbarry/CollabSource/back-end/errors"
	// "github.com/bfbarry/CollabSource/back-end/controllers"
)

func ResponseEntity(w http.ResponseWriter, status int, Body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(Body)
}
