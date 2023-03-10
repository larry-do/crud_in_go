package router

import (
	"encoding/json"
	"net/http"
)

func (resp Response) RespondJson(object any) {
	resp.ResponseWriter.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(resp).Encode(object)
	if err != nil {
		return
	}
}

type Response struct {
	http.ResponseWriter
}
