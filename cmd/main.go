package main

import (
	"flag"
	"fmt"
	"log"
	"spamChecker/config"
	"spamChecker/database"
	"spamChecker/middleware"
	"spamChecker/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	mode := flag.String("mode", "development", "Application mode: development | production")
	flag.Parse()

	if *mode == "production" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("Starting in PRODUCTION mode")
	} else {
		gin.SetMode(gin.DebugMode)
		fmt.Println("Starting in DEVELOPMENT mode")
	}

	config.LoadConfig(*mode)
	database.Connect()

	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	routes.SetupRoutes(r)

	log.Fatal(r.Run(":8080"))
}
