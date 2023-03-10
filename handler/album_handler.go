package handler

import (
	"CRUD/datasource"
	"CRUD/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteAlbum(resp http.ResponseWriter, req *http.Request) {
	var album model.Album
	id := mux.Vars(req)["id"]
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		_ = json.NewEncoder(resp).Encode("album not found!")
		return
	}
	datasource.Connection().Delete(&album, id)
	resp.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(resp).Encode("album deleted successfully")
}

func UpdateAlbum(resp http.ResponseWriter, req *http.Request) {
	var album model.Album
	id := mux.Vars(req)["id"]
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		_ = json.NewEncoder(resp).Encode("album not found!")
		return
	}
	_ = json.NewDecoder(req.Body).Decode(&album)
	datasource.Connection().Save(&album)
	resp.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(resp).Encode(album)
}

func GetAlbums(resp http.ResponseWriter, req *http.Request) {
	var albums []*model.Album
	datasource.Connection().Find(&albums)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(resp).Encode(albums)
}

func GetAlbum(resp http.ResponseWriter, req *http.Request) {
	var album model.Album
	id := mux.Vars(req)["id"]
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		_ = json.NewEncoder(resp).Encode("album not found!")
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(resp).Encode(album)
}

func PostAlbum(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var album model.Album
	_ = json.NewDecoder(req.Body).Decode(&album)
	datasource.Connection().Create(&album)
	_ = json.NewEncoder(resp).Encode(album)
}
