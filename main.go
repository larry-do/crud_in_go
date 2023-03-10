package main

import (
	"CRUD/datasource"
	"CRUD/handler"
	"CRUD/router"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

func main() {
	datasource.Initialize()

	var appRouter = router.Router{
		Router: mux.NewRouter(),
	}

	appRouter.HandleGetRequest("/", func(writer http.ResponseWriter, request *http.Request) {
		// do nothing
	})

	appRouter.HandlePostRequest("/api/v1/album", handler.PostAlbum)
	appRouter.HandleGetRequest("/api/v1/albums", handler.GetAlbums)
	appRouter.HandleGetRequest("/api/v1/album/{id}", handler.GetAlbum)
	appRouter.HandlePutRequest("/api/v1/album/{id}", handler.UpdateAlbum)
	appRouter.HandleDeleteRequest("/api/v1/album/{id}", handler.DeleteAlbum)

	appRouter.HandlePostRequest("/api/v1/artist", handler.PostArtist)
	appRouter.HandleGetRequest("/api/v1/artists", handler.GetArtists)
	appRouter.HandleGetRequest("/api/v1/artist/{id}", handler.GetArtist)
	appRouter.HandlePutRequest("/api/v1/artist/{id}", handler.UpdateArtist)
	appRouter.HandleDeleteRequest("/api/v1/artist/{id}", handler.DeleteArtist)

	fmt.Println("Listening and serving on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", appRouter.Router))
}
