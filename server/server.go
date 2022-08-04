package main

import (
	"net/http"

	"github.com/kimcodell/SH-MH-diary/server/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "SH ❤️ MH Diary Server")
	})

	routers.ConnectPostRouter(r)

	r.Run(":8000")
}
