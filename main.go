package main

import (
	"example/web-service-gin/configure"
	"example/web-service-gin/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	// Load the env here before calling
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Configure the database connection
	configure.ConfigureDB()

	// Set up the Gin router
	server := os.Getenv("SERVER_ADDRESS")
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-type"},
		AllowAllOrigins: true,
	}))

	// Set up the routes
	routes.SetRoutes(router)

	// Start the server
	logrus.Error(router.Run(server))
}
