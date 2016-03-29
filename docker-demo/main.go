package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "INFO_PATH": c.Request.URL.Path,
            "HTTP_STATUS": http.StatusOK,
        })
    })
    r.Run(":8001") // listen and server on 0.0.0.0:8001
}
