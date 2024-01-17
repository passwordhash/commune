package main

import (
	"commune/internal/handler"
	"commune/internal/repository"
	"commune/internal/service"
	"commune/internal/websocket"
	pkgRepo "commune/pkg/repository"
	"commune/pkg/server"
	"flag"
	"github.com/sirupsen/logrus"
)

//var ctx = context.TODO()

func main() {
	flag.Parse()

	hub := websocket.NewHub()
	go hub.Run()

	mongoDB, err := pkgRepo.NewMongoDB(pkgRepo.ConnectURI)
	if err != nil {
		logrus.Fatalln("connection to mongo db failed")
	}

	repos := repository.NewRepository(mongoDB)
	services := service.NewService(repos)
	handlers := handler.NewHandler(hub, services)

	srv := new(server.Server)
	if err := srv.Run("8090", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	logrus.Info("Commune Server Started")

}
