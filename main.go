package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    app := gin.Default()
    app.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World V3!",
        })
    })
    app.Run(":8080")
}
