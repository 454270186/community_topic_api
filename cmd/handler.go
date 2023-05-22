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
	var newTopic dto.NewTopicReq
	if err := c.ShouldBind(&newTopic); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data := ctrl.AddNewTopic(newTopic)
	c.JSON(http.StatusOK, data)
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

func DeleteTopicById(c *gin.Context) {
	topicId := c.Param("id")

	data := ctrl.DeleteTopic(topicId)
	c.JSON(200, data)
}

func DeletePostById(c *gin.Context) {
	postId := c.Param("id")

	data := ctrl.DeletePost(postId)
	c.JSON(200, data)
}