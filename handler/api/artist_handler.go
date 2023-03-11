package api

import (
	"CRUD/datasource"
	"CRUD/handler"
	"CRUD/model"
	"CRUD/router"
	"net/http"
)

func DeleteArtist(resonse router.Response, request router.Request) {
	id := handler.PathVariable(request.Request, "id")
	var artist model.Artist
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		handler.RespondJson(resonse, "artist not found!")
		return
	}
	datasource.Connection().Delete(&artist, id)
	handler.SetContentTypeJson(resonse)
	handler.RespondJson(resonse, "artist deleted successfully")
}

func UpdateArtist(resonse router.Response, request router.Request) {
	id := handler.PathVariable(request.Request, "id")
	var artist model.Artist
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		handler.RespondJson(resonse, "artist not found!")
		return
	}
	handler.RequestBody(request.Request, &artist)
	datasource.Connection().Save(&artist)
	handler.SetContentTypeJson(resonse)
	handler.RespondJson(resonse, artist)
}

func GetArtists(resonse router.Response, request router.Request) {
	var artists []*model.Artist
	datasource.Connection().Find(&artists)
	handler.SetContentTypeJson(resonse)
	resonse.WriteHeader(http.StatusOK)
	handler.RespondJson(resonse, artists)
}

func GetArtist(resonse router.Response, request router.Request) {
	id := handler.PathVariable(request.Request, "id")
	var artist model.Artist
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		handler.RespondJson(resonse, "artist not found!")
		return
	}
	handler.SetContentTypeJson(resonse)
	handler.RespondJson(resonse, artist)
}

func PostArtist(resonse router.Response, request router.Request) {
	handler.SetContentTypeJson(resonse)
	var artist model.Artist
	handler.RequestBody(request.Request, &artist)
	datasource.Connection().Create(&artist)
	handler.RespondJson(resonse, artist)
}
