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
