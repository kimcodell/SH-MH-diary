package services

import (
	"net/http"
	"strconv"

	"github.com/kimcodell/SH-MH-diary/server/repositories"
	"github.com/kimcodell/SH-MH-diary/server/types"
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
	var params types.PostCreateDto
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	isSuccess := repositories.CreatePost(params)

	statusCode := http.StatusOK
	if !isSuccess {
		statusCode = http.StatusInternalServerError
	}
	c.JSON(statusCode, gin.H{"success": isSuccess})
}

func UpdatePostById(c *gin.Context) {
	var params types.PostCreateDto
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isSuccess := repositories.UpdatePost(params)

	statusCode := http.StatusOK
	if !isSuccess {
		statusCode = http.StatusInternalServerError
	}
	c.JSON(statusCode, gin.H{"success": isSuccess})
}

func DeletePostById(c *gin.Context) {

}
