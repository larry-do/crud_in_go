package handler

import (
	"CRUD/datasource"
	"CRUD/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteArtist(resp http.ResponseWriter, req *http.Request) {
	var artist model.Artist
	id := mux.Vars(req)["id"]
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		json.NewEncoder(resp).Encode("artist not found!")
		return
	}
	datasource.Connection().Delete(&artist, id)
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode("artist deleted successfully")
}

func UpdateArtist(resp http.ResponseWriter, req *http.Request) {
	var artist model.Artist
	id := mux.Vars(req)["id"]
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		json.NewEncoder(resp).Encode("artist not found!")
		return
	}
	json.NewDecoder(req.Body).Decode(&artist)
	datasource.Connection().Save(&artist)
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(artist)
}

func GetArtists(resp http.ResponseWriter, req *http.Request) {
	var artists []*model.Artist
	datasource.Connection().Find(&artists)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(artists)
}

func GetArtist(resp http.ResponseWriter, req *http.Request) {
	var artist model.Artist
	id := mux.Vars(req)["id"]
	datasource.Connection().First(&artist, id)
	if artist.ID == 0 {
		json.NewEncoder(resp).Encode("artist not found!")
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(artist)
}

func PostArtist(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var artist model.Artist
	json.NewDecoder(req.Body).Decode(&artist)
	datasource.Connection().Create(&artist)
	json.NewEncoder(resp).Encode(artist)
}
