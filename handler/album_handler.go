package handler

import (
	"CRUD/datasource"
	"CRUD/model"
	"net/http"
)

func DeleteAlbum(respWriter http.ResponseWriter, request *http.Request) {
	var resp, req = convert(respWriter, request)
	var id = req.PathVariable("id")
	var album model.Album
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		resp.RespondJson("album not found!")
		return
	}
	datasource.Connection().Delete(&album, id)
	resp.RespondJson("album deleted successfully")
}

func UpdateAlbum(respWriter http.ResponseWriter, request *http.Request) {
	var resp, req = convert(respWriter, request)
	var id = req.PathVariable("id")
	var album model.Album
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		resp.RespondJson("album not found!")
		return
	}
	req.RequestBodyFromJson(&album)
	datasource.Connection().Save(&album)
	resp.RespondJson(album)
}

func GetAlbums(respWriter http.ResponseWriter, request *http.Request) {
	var resp, _ = convert(respWriter, request)
	var albums []*model.Album
	datasource.Connection().Find(&albums)
	resp.WriteHeader(http.StatusOK)
	resp.RespondJson(albums)
}

func GetAlbum(respWriter http.ResponseWriter, request *http.Request) {
	var resp, req = convert(respWriter, request)
	var id = req.PathVariable("id")

	var album model.Album
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		resp.RespondJson("album not found!")
		return
	}
	resp.RespondJson(album)
}

func PostAlbum(respWriter http.ResponseWriter, request *http.Request) {
	var resp, req = convert(respWriter, request)
	var album = req.GetRequestBodyFromJson()
	datasource.Connection().Create(&album)
	resp.RespondJson(album)
}
