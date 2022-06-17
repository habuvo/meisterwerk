package main

import (
	"meisterwerk/handlers"
	"meisterwerk/repositories/postgres"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if err := godotenv.Load(".env"); err != nil {
		logger.Fatal("can't read config", zap.Error(err))
	}

	pRepo, err := postgres.Connect()
	if err != nil {
		logger.Fatal("can't open DB connect", zap.Error(err))
	}
	defer pRepo.Close()

	router := gin.Default()
	eventController := handlers.NewEventer(pRepo)

	registerRoutes(router, eventController)

	router.Run()
}

func registerRoutes(router *gin.Engine, events handlers.Eventer) {
	router.GET("/events", events.Get())
}
