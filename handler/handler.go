package handler

import (
	"CRUD/router"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func SetContentTypeJson(resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
}

func PathVariable(req *http.Request, key string) string {
	return mux.Vars(req)[key]
}

func RespondJson(resp http.ResponseWriter, object any) {
	json.NewEncoder(resp).Encode(object)
}

func RequestBody(req *http.Request, object any) any {
	return json.NewDecoder(req.Body).Decode(object)
}

func Convert(resp http.ResponseWriter, req *http.Request) (router.Response, router.Request) {
	return router.Response{
			ResponseWriter: resp,
		}, router.Request{
			Request: req,
		}
}
