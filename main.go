package main

import (
	"CRUD/datasource"
	"CRUD/handler"
	"CRUD/handler/api"
	"CRUD/router"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

func main() {
	datasource.Initialize()

	// keep this statement last
	initializeRoutes()
}

func initializeRoutes() {
	var appRouter = router.Router{
		Router: mux.NewRouter(),
	}

	appRouter.HandleGetRequest("/", func(writer http.ResponseWriter, request *http.Request) {
		// do nothing
	})

	appRouter.HandlePostRequest("/api/v1/album", api.PostAlbum)
	appRouter.HandleGetRequest("/api/v1/albums", api.GetAlbums)
	appRouter.HandleGetRequest("/api/v1/album/{id}", api.GetAlbum)
	appRouter.HandlePutRequest("/api/v1/album/{id}", api.UpdateAlbum)
	appRouter.HandleDeleteRequest("/api/v1/album/{id}", api.DeleteAlbum)

	appRouter.HandlePostRequest("/api/v1/artist", api.PostArtist)
	appRouter.HandleGetRequest("/api/v1/artists", api.GetArtists)
	appRouter.HandleGetRequest("/api/v1/artist/{id}", api.GetArtist)
	appRouter.HandlePutRequest("/api/v1/artist/{id}", api.UpdateArtist)
	appRouter.HandleDeleteRequest("/api/v1/artist/{id}", api.DeleteArtist)

	appRouter.HandleGetRequest("/artist/{id}", handler.GetArtist)

	log.Println("Listening and serving on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", appRouter.Router))
}
