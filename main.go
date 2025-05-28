package main

import (
	"ADMISSION-MANAGEMENT-SYSTEM/config"
	"ADMISSION-MANAGEMENT-SYSTEM/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run() // default on port 8080
}
