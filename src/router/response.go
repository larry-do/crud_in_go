package router

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func (resp Response) RespondJson(object any) {
	resp.ResponseWriter.Header().Set("Content-Type", "application/json")
	var err = json.NewEncoder(resp).Encode(object)
	if err != nil {
		return
	}
}

func (resp Response) ResponseCode(statusCode int) Response {
	resp.WriteHeader(statusCode)
	return resp
}

func (resp Response) RespondHtmlView(viewTemplate string, model any) (Response, error) {
	var tmpl = template.Must(template.ParseFiles(viewTemplate))
	var err = tmpl.Execute(resp, model)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type Response struct {
	http.ResponseWriter
}
