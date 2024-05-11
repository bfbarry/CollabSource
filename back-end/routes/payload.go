package routes

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/bfbarry/CollabSource/back-end/errors"
	"github.com/bfbarry/CollabSource/back-end/controllers"
)


func writeJsonSuccess(w http.ResponseWriter, res *controllers.Resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	w.Write(res.Body)
}

func writeJsonError(w http.ResponseWriter, e *errors.Error) {
	//TODO: if status >= 500, also handle internally, else write to response
	log.Printf("%d: %s", e.Status(), e.Error())
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