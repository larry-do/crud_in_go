package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (req Request) PathVariable(key string) any {
	return mux.Vars(req.Request)[key]
}

func (req Request) RequestBodyFromJson(object any) any {
	return json.NewDecoder(req.Body).Decode(object)
}

func (req Request) GetRequestBodyFromJson() any {
	var object any
	err := json.NewDecoder(req.Body).Decode(&object)
	if err != nil {
		return nil
	}
	return object
}

type Request struct {
	*http.Request
}
