package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"go-personal-website-api/configs"
	httpHandler "go-personal-website-api/services/handler/http"
	"go-personal-website-api/services/repository"
	"go-personal-website-api/services/usecase"
)

func main() {
	// Load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	var cfg configs.MainConfig
	assignConfig(&cfg)

	// log.Printf("Check config sugarbae06: %+v", cfg)

	// Initiate Database
	dbConn, err := configs.InitiateDatabase(cfg.DatabaseConfig)
	if err != nil {
		log.Fatal("Error initiate database", err)
	}

	// Initiate Repository
	repo := repository.NewRepository(dbConn)

	// Initiate Usecase
	serviceUc := usecase.NewUsecase(repo)

	// Initiate Router
	e := echo.New()
	address := fmt.Sprintf(":%s", cfg.AppPort)
	httpHandler.NewHttpHandler(e, serviceUc)
	e.Logger.Fatal(e.Start(address))
}

func assignConfig(cfg *configs.MainConfig) {
	// Apps
	cfg.AppPort = os.Getenv("APP_PORT")

	// Database
	cfg.DatabaseConfig.DBEngine = os.Getenv("DB_ENGINE")
	cfg.DatabaseConfig.DBHost = os.Getenv("DB_HOST")
	cfg.DatabaseConfig.DBPort = os.Getenv("DB_PORT")
	cfg.DatabaseConfig.DBUser = os.Getenv("DB_USER")
	cfg.DatabaseConfig.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DatabaseConfig.DBName = os.Getenv("DB_NAME")
}
