package handler

import (
	"CRUD/datasource"
	"CRUD/model"
	"net/http"
)

func DeleteArtist(resp http.ResponseWriter, req *http.Request) {
	id := pathVariable(req, "id")
	var artist model.Artist
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		respondJson(resp, "artist not found!")
		return
	}
	datasource.Connection().Delete(&artist, id)
	setContentTypeJson(resp)
	respondJson(resp, "artist deleted successfully")
}

func UpdateArtist(resp http.ResponseWriter, req *http.Request) {
	id := pathVariable(req, "id")
	var artist model.Artist
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		respondJson(resp, "artist not found!")
		return
	}
	requestBody(req, &artist)
	datasource.Connection().Save(&artist)
	setContentTypeJson(resp)
	respondJson(resp, artist)
}

func GetArtists(resp http.ResponseWriter, req *http.Request) {
	var artists []*model.Artist
	datasource.Connection().Find(&artists)
	setContentTypeJson(resp)
	resp.WriteHeader(http.StatusOK)
	respondJson(resp, artists)
}

func GetArtist(resp http.ResponseWriter, req *http.Request) {
	id := pathVariable(req, "id")
	var artist model.Artist
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		respondJson(resp, "artist not found!")
		return
	}
	setContentTypeJson(resp)
	respondJson(resp, artist)
}

func PostArtist(resp http.ResponseWriter, req *http.Request) {
	setContentTypeJson(resp)
	var artist model.Artist
	requestBody(req, &artist)
	datasource.Connection().Create(&artist)
	respondJson(resp, artist)
}
