package handler

import (
	"CRUD/src/datasource"
	"CRUD/src/model"
	"CRUD/src/router"
	"net/http"
)

func GetArtist(resonse router.Response, request router.Request) {
	var id = request.PathVariable("id")
	var artist model.Artist
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		resonse.ResponseCode(http.StatusNotFound)
		return
	}
	resonse.RespondHtmlView("resources/template/artist.html", artist)
}

func HandleCreateArtistView(resonse router.Response, request router.Request) {
	resonse.RespondHtmlView("resources/template/create_artist.html", nil)
}

func CreateArtist(resonse router.Response, request router.Request) {
	if request.FormValue("name") == "" {
		resonse.RespondHtmlView("resources/template/create_artist.html", struct{ Success bool }{false})
		return
	}

	var artist = model.Artist{
		Name: request.FormValue("name"),
		// Description: req.FormValue("description"),
	}

	datasource.Connection().Save(&artist)
	resonse.RespondHtmlView("resources/template/create_artist.html", struct{ Success bool }{true})
}
