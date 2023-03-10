package router

import "net/http"
import "github.com/gorilla/mux"

func (router Router) HandlePostRequest(path string, doFunc func(http.ResponseWriter,
	*http.Request)) *mux.Route {
	return router.HandleFunc(path, doFunc).Methods("POST")
}

func (router Router) HandleGetRequest(path string, doFunc func(http.ResponseWriter,
	*http.Request)) *mux.Route {
	return router.HandleFunc(path, doFunc).Methods("GET")
}

func (router Router) HandlePutRequest(path string, doFunc func(http.ResponseWriter,
	*http.Request)) *mux.Route {
	return router.HandleFunc(path, doFunc).Methods("PUT")
}

func (router Router) HandleDeleteRequest(path string, doFunc func(http.ResponseWriter,
	*http.Request)) *mux.Route {
	return router.HandleFunc(path, doFunc).Methods("DELETE")
}

type Router struct {
	*mux.Router
}
