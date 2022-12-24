package main

import (
    "net/http"
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/hello", func(c *gin.Context) {
        // Create a new user
        // (omitted for brevity)
        c.JSON(http.StatusOK, "Hello, World!")
        fmt.Println("Hello, World!")
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}
