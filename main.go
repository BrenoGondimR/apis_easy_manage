package main

import (
	"example/web-service-gin/routes"
	"github.com/gin-contrib/cors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// load the env here before calling
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := os.Getenv("SERVER_ADDRESS")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-type"},
		AllowAllOrigins: true,
	}))
	routes.SetRoutes(router)

	logrus.Error(router.Run(server))
}
