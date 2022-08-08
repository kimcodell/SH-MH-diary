package main

import (
	"net/http"

	"github.com/kimcodell/SH-MH-diary/server/database"
	"github.com/kimcodell/SH-MH-diary/server/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectToDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "SH ❤️ MH Diary Server")
	})
	
	routers.ConnectPostRouter(r)

	r.Run(":8000")
}
