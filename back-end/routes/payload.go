package routes

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/bfbarry/CollabSource/back-end/errors"
)


func writeJsonSuccess(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func writeJsonError(w http.ResponseWriter, e *errors.Error) {
	//TODO: if status >= 500, also handle internally, else write to response
	log.Print(e.Error())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status())
	msg := e.ClientMessage()
	res, err := json.Marshal(msg)

	if err != nil { // TODO: 500
		log.Printf("Could not marshal json in writeJsonError, %s", err.Error())
		return
	}
	w.Write(res)
}