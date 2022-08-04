package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}