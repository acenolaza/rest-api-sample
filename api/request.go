package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Request struct {
	Request *http.Request
}

func (r *Request) GetJSONBody(model interface{}) error {
	if err := json.NewDecoder(r.Request.Body).Decode(&model); err != nil {
		log.Printf("could not deserialize body into %T object", model)
		return err
	}
	return nil
}

func (r *Request) GetVarID() (int, error) {
	vars := mux.Vars(r.Request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("could not convert parameter id to int")
		return 0, err
	}
	return id, nil
}
