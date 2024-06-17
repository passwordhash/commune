package main

import (
	"commune/internal/handler"
	"commune/internal/repository"
	"commune/internal/service"
	"commune/internal/websocket"
	"commune/pkg/app"
	pkgRepo "commune/pkg/repository"
	"commune/pkg/server"
	"flag"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

//var ctx = context.TODO()

// ConfigFilename TEMP
var ConfigFilename = "config"

func main() {
	flag.Parse()

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("cannot load env file: %v", err.Error())
	}

	if err := app.InitConfig(ConfigFilename); err != nil {
		logrus.Fatal(err)
	}
	config := LoadConfig()

	hub := websocket.NewHub()
	go hub.Run()

	mongoDB, err := pkgRepo.NewMongoDB(pkgRepo.ConnectURI)
	if err != nil {
		logrus.Fatalln("connection to mongo db failed")
	}

	repos := repository.NewRepository(mongoDB)
	services := service.NewService(repos, service.Deps{
		AccessTokenTTL: viper.GetDuration("auth.accessTokenTTL"),
		SigingKey:      os.Getenv("SIGING_KEY"),
		PassphraseSalt: os.Getenv("PASSWORD_SALT"),
	})
	handlers := handler.NewHandler(hub, services)

	srv := new(server.Server)
	logrus.Info("Trying to start Commune Server...")
	if err := srv.Run(config.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	logrus.Info("Commune Server shut down")
}

type Config struct {
	Port string
}

func LoadConfig() Config {
	port := viper.GetString("http.port")
	return Config{Port: port}
}
