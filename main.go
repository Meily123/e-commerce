package main

import (
	"WebAPI/config"
	"WebAPI/docs"
	"WebAPI/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

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

	//swagger info
	docs.SwaggerInfo.Title = "Swagger E-commerce API"
	docs.SwaggerInfo.Description = "e-commerce documentation."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", port)
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// router
	router := gin.Default()

	// group version router
	versionRouter := router.Group("/v1")

	// routes
	route.ProductRoute(versionRouter)

	// run app
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("error running app")
	}

}
