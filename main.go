package main

import (
	"WebAPI/config"
	"WebAPI/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
	}

	// connect and migrate database
	config.ConnectToDatabase()

	// get port from env
	port := os.Getenv("PORT")

	// router
	router := gin.Default()

	// group version router
	versionRouter := router.Group("/v1")

	// routes
	route.ProductRoute(versionRouter)

	// run app
	err = router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("error running app")
	}

}
