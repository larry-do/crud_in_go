package api

import (
	"CRUD/datasource"
	"CRUD/model"
	"CRUD/router"
	"net/http"
)

func DeleteAlbum(resonse router.Response, request router.Request) {
	var id = request.PathVariable("id")
	var album model.Album
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		resonse.RespondJson("album not found!")
		return
	}
	datasource.Connection().Delete(&album, id)
	resonse.RespondJson("album deleted successfully")
}

func UpdateAlbum(resonse router.Response, request router.Request) {
	var id = request.PathVariable("id")
	var album model.Album
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		resonse.RespondJson("album not found!")
		return
	}
	request.RequestBodyFromJson(&album)
	datasource.Connection().Save(&album)
	resonse.RespondJson(album)
}

func GetAlbums(resonse router.Response, request router.Request) {
	var albums []*model.Album
	datasource.Connection().Find(&albums)
	resonse.WriteHeader(http.StatusOK)
	resonse.RespondJson(albums)
}

func GetAlbum(resonse router.Response, request router.Request) {
	var id = request.PathVariable("id")

	var album model.Album
	datasource.Connection().First(&album, id)
	if album.ID == 0 {
		resonse.RespondJson("album not found!")
		return
	}
	resonse.RespondJson(album)
}

func PostAlbum(resonse router.Response, request router.Request) {
	var album = request.GetRequestBodyFromJson()
	datasource.Connection().Create(&album)
	resonse.RespondJson(album)
}
