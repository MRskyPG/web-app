package main

import (
	"github.com/MRskyPG/web-app"
	"github.com/MRskyPG/web-app/pkg/handler"
	"log"
)

func main() {
	var handlers handler.Handler
	var srv web.Server

	if err := srv.Run("80", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}

}
