package main

import (
	"context"
	"github.com/MRskyPG/web-app"
	"github.com/MRskyPG/web-app/pkg/handler"
	"github.com/MRskyPG/web-app/pkg/service"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mapService := service.NewMapService()
	handlers := handler.New(mapService)
	var srv web.Server

	//Using graceful shutdown
	go func() {
		if err := srv.Run("80", handlers.InitRoutes()); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Web-app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Web-app Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error was occured on server shutting down: %s", err.Error())
	}
}
