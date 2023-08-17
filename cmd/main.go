package main

import (
	"github.com/MRskyPG/web-app"
	"github.com/MRskyPG/web-app/pkg/handler"
	"github.com/MRskyPG/web-app/pkg/service"
	"log"
)

func main() {
	mapService := service.NewMapService()
	handlers := handler.New(mapService)
	var srv web.Server

	if err := srv.Run("80", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}

}
