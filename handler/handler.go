package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// SetContentTypeJson deprecated
func SetContentTypeJson(resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
}

// PathVariable deprecated
func PathVariable(req *http.Request, key string) string {
	return mux.Vars(req)[key]
}

// RespondJson deprecated
func RespondJson(resp http.ResponseWriter, object any) {
	json.NewEncoder(resp).Encode(object)
}

// RequestBody deprecated
func RequestBody(req *http.Request, object any) any {
	return json.NewDecoder(req.Body).Decode(object)
}
