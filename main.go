package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    app := gin.Default()
    app.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World Release V1",
        })
    })
    app.Run(":8080")
}
