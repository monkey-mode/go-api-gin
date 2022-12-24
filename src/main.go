package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
)

// @title Pratice Go API
// @version 1.0
// @description.markdown
// @termsOfService http://somewhere.com/

// @contact.name API Support
// @contact.url http://somewhere.com/support
// @contact.email support@somewhere.com

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := gin.Default()
    // Use the `swaggo/gin-swagger` middleware to generate Swagger documentation for the API endpoints.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Hello godoc
	// @summary Helllo World
	// @description Hello world for the service
	// @id HelloWorld
	// @produce plain
	// @response 200 {string} string "Hello, World!"
	// @router /hello [get]
	r.GET("/hello", func(c *gin.Context) {
		// Create a new user
		// (omitted for brevity)
		c.String(http.StatusOK, "Hello, World!")
		fmt.Println("Hello, World!")
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
