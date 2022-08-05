package services

import (
	"net/http"
	"strconv"

	"github.com/kimcodell/SH-MH-diary/server/repositories"
	"github.com/kimcodell/SH-MH-diary/server/utils"

	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	posts := repositories.GetAllPosts()
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	intId, convertErr := strconv.Atoi(id)
	utils.CatchError(utils.ErrorParams{Err: convertErr, Message: "Post id is not valid."})

	post := repositories.GetPostById(intId)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func CreatePost(c *gin.Context) {

	repositories.CreatePost()
	c.JSON(http.StatusOK, gin.H{"success": true})
}
