package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func setContentTypeJson(resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
}

func pathVariable(req *http.Request, pathVariableName string) string {
	return mux.Vars(req)[pathVariableName]
}

func respondJson(resp http.ResponseWriter, object any) {
	json.NewEncoder(resp).Encode(object)
}

func requestBody(req *http.Request, object any) any {
	return json.NewDecoder(req.Body).Decode(object)
}
