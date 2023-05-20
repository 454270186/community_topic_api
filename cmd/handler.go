package main

import (
	"net/http"

	"github.com/454270186/CommuTopicPage/controller"
	"github.com/454270186/CommuTopicPage/dto"
	"github.com/gin-gonic/gin"
)

var ctrl controller.Controller

// [GET] /community/page/:id
func GetTopicById(c *gin.Context) {
	topicId := c.Param("id")
	data := ctrl.QueryPageInfo(topicId)
	c.JSON(http.StatusOK, data)
}

func CreateNewTopic(c *gin.Context) {

}

func CreateNewPostById(c *gin.Context) {
	// get post body
	var newPost dto.NewPostReq
	if err := c.ShouldBind(&newPost); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data := ctrl.AddNewPost(newPost)
	c.JSON(http.StatusOK, data)
}