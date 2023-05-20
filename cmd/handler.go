package main

import (
	"net/http"

	"github.com/454270186/CommuTopicPage/controller"
	"github.com/gin-gonic/gin"
)

var ctrl controller.Controller

// GET /community/page/:id
func GetTopicById(c *gin.Context) {
	topicId := c.Param("id")	
	data := ctrl.QueryPageInfo(topicId)
	c.JSON(http.StatusOK, data)
}