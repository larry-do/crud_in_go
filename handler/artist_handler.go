package handler

import (
	"CRUD/datasource"
	"CRUD/model"
	"net/http"
)

func GetArtist(respWriter http.ResponseWriter, request *http.Request) {
	var resp, req = Convert(respWriter, request)
	var id = req.PathVariable("id")
	var artist model.Artist
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		resp.ResponseCode(http.StatusNotFound)
		return
	}
	resp.RespondHtmlView("template/artist.html", artist)
}

func HandleCreateArtistView(respWriter http.ResponseWriter, request *http.Request) {
	var resp, _ = Convert(respWriter, request)
	resp.RespondHtmlView("template/create_artist.html", nil)
}

func CreateArtist(respWriter http.ResponseWriter, request *http.Request) {
	var resp, req = Convert(respWriter, request)

	if req.FormValue("name") == "" {
		resp.RespondHtmlView("template/create_artist.html", struct{ Success bool }{false})
		return
	}

	var artist = model.Artist{
		Name: req.FormValue("name"),
		// Description: req.FormValue("description"),
	}

	datasource.Connection().Save(&artist)
	resp.RespondHtmlView("template/create_artist.html", struct{ Success bool }{true})
}
