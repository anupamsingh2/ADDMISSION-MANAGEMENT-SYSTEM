package main

import (
	"ADMISSION-MANAGEMENT-SYSTEM/config"
	"ADMISSION-MANAGEMENT-SYSTEM/routes"
	"ADMISSION-MANAGEMENT-SYSTEM/utils"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Debug print to verify JWT_SECRET loaded correctly
	fmt.Println("JWT_SECRET from env:", os.Getenv("JWT_SECRET"))

	// Initialize JWT secret AFTER loading .env
	utils.InitializeJWT()

	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run() // defaults to ":8080"
}
