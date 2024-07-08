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
	"time"
)

//var ctx = context.TODO()

// ConfigFilename TEMP
var ConfigFilename = "config"

// TODO:
// 1) список сообщений с информацией об авторе

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

	mongoDB, err := pkgRepo.NewMongoDB(os.Getenv("MONGO_URI"))
	if err != nil {
		logrus.Fatalln("connection to mongo db failed")
	}

	repos := repository.NewRepository(mongoDB)
	services := service.NewService(repos, service.Deps{
		AccessTokenTTL: config.AccessTokenTTL,
		SigingKey:      os.Getenv("SIGING_KEY"),
		PasscodeSalt:   os.Getenv("PASSWORD_SALT"),
		EmailDeps: service.EmailDeps{
			SmtpHost: config.smtpHost,
			SmtpPort: config.smtpPort,
			From:     config.from,
			Password: os.Getenv("SMTP_PASSWORD"),
		},
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
	AccessTokenTTL time.Duration

	Port string

	// For email
	smtpHost string
	smtpPort int
	from     string
}

func LoadConfig() Config {
	accessTokenTTl := viper.GetDuration("auth.accessTokenTTL")
	port := viper.GetString("http.port")

	stmpHost := viper.GetString("email.smtpHost")
	smtpPort := viper.GetInt("email.smtpPort")
	from := viper.GetString("email.from")
	return Config{
		AccessTokenTTL: accessTokenTTl,
		Port:           port,
		smtpPort:       smtpPort,
		smtpHost:       stmpHost,
		from:           from,
	}
}
