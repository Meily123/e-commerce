package main

import (
	"WebAPI/config"
	"WebAPI/docs"
	"WebAPI/route"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func init() {
	config.LoadEnvVariable()
	config.ConnectToDatabase()
}

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

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

	// routeInit
	route.InitRoute(router, "/v1")

	// run app
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("error running app")
	}

}
