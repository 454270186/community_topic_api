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
		
		// post like
		comPage.PUT("/post/:postid/like", AddPostLikeCnt)
		comPage.GET("/post/:topicid/like", GetPostByLike)
	}

	return router
}
