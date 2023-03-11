package api

import (
	"CRUD/datasource"
	"CRUD/handler"
	"CRUD/model"
	"net/http"
)

func DeleteArtist(resp http.ResponseWriter, req *http.Request) {
	id := handler.PathVariable(req, "id")
	var artist model.Artist
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		handler.RespondJson(resp, "artist not found!")
		return
	}
	datasource.Connection().Delete(&artist, id)
	handler.SetContentTypeJson(resp)
	handler.RespondJson(resp, "artist deleted successfully")
}

func UpdateArtist(resp http.ResponseWriter, req *http.Request) {
	id := handler.PathVariable(req, "id")
	var artist model.Artist
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		handler.RespondJson(resp, "artist not found!")
		return
	}
	handler.RequestBody(req, &artist)
	datasource.Connection().Save(&artist)
	handler.SetContentTypeJson(resp)
	handler.RespondJson(resp, artist)
}

func GetArtists(resp http.ResponseWriter, req *http.Request) {
	var artists []*model.Artist
	datasource.Connection().Find(&artists)
	handler.SetContentTypeJson(resp)
	resp.WriteHeader(http.StatusOK)
	handler.RespondJson(resp, artists)
}

func GetArtist(resp http.ResponseWriter, req *http.Request) {
	id := handler.PathVariable(req, "id")
	var artist model.Artist
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		handler.RespondJson(resp, "artist not found!")
		return
	}
	handler.SetContentTypeJson(resp)
	handler.RespondJson(resp, artist)
}

func PostArtist(resp http.ResponseWriter, req *http.Request) {
	handler.SetContentTypeJson(resp)
	var artist model.Artist
	handler.RequestBody(req, &artist)
	datasource.Connection().Create(&artist)
	handler.RespondJson(resp, artist)
}
