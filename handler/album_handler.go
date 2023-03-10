package handler

import (
	"CRUD/datasource"
	"CRUD/model"
	"net/http"
)

func DeleteAlbum(resp http.ResponseWriter, req *http.Request) {
	id := pathVariable(req, "id")
	var album model.Album
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		respondJson(resp, "album not found!")
		return
	}
	datasource.Connection().Delete(&album, id)
	setContentTypeJson(resp)
	respondJson(resp, "album deleted successfully")
}

func UpdateAlbum(resp http.ResponseWriter, req *http.Request) {
	id := pathVariable(req, "id")
	var album model.Album
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		respondJson(resp, "album not found!")
		return
	}
	requestBody(req, &album)
	datasource.Connection().Save(&album)
	setContentTypeJson(resp)
	respondJson(resp, album)
}

func GetAlbums(resp http.ResponseWriter, req *http.Request) {
	var albums []*model.Album
	datasource.Connection().Find(&albums)
	setContentTypeJson(resp)
	resp.WriteHeader(http.StatusOK)
	respondJson(resp, albums)
}

func GetAlbum(resp http.ResponseWriter, req *http.Request) {
	id := pathVariable(req, "id")
	var album model.Album
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		respondJson(resp, "album not found!")
		return
	}
	setContentTypeJson(resp)
	respondJson(resp, album)
}

func PostAlbum(resp http.ResponseWriter, req *http.Request) {
	setContentTypeJson(resp)
	var album model.Album
	requestBody(req, &album)
	datasource.Connection().Create(&album)
	respondJson(resp, album)
}
