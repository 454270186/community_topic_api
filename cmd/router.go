package main

import (
	"github.com/454270186/CommuTopicPage/controller"
	"github.com/454270186/CommuTopicPage/repository"
	"github.com/454270186/CommuTopicPage/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ctrl = controller.NewController(service.NewPageService(repository.NewDataRepo(db), rdb))
	router := gin.Default()

	comPage := router.Group("/community/page")
	{
		comPage.GET("/:id", GetTopicById)
		comPage.POST("/topic", CreateNewTopic)
		comPage.POST("/post", CreateNewPostById)
		comPage.DELETE("/topic/:id", DeleteTopicById)
		comPage.DELETE("/post/:id", DeletePostById)
		
		comPage.PUT("/post/:postid/like", AddPostLikeCnt)
	}

	return router
}
