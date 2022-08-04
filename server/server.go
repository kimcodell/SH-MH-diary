package main

import (
	"net/http"

	"github.com/kimcodell/SH-MH-diary/server/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
	
	routers.ConnectPostRouter(r)

  r.Run()
}