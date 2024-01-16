package main

import (
	"commune"
	"commune/internal/handler"
	"commune/pkg/server"
	"flag"
	"github.com/sirupsen/logrus"
)

//var ctx = context.TODO()

func main() {
	flag.Parse()
	//mongo, err := repository.NewMongoDB()
	//if err != nil {
	//	log.Fatal("mongo err: ", err)
	//}
	//defer mongo.Disconnect(ctx)
	//
	//db := mongo.Database("commune")
	//_ = db.Collection("users")

	hub := commune.NewHub()
	go hub.Run()

	handlers := &handler.Handler{
		Hub: hub,
	}

	srv := new(server.Server)
	if err := srv.Run("8090", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	logrus.Info("Commune Server Started")

}
