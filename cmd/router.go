package main

import (
	"github.com/454270186/CommuTopicPage/controller"
	"github.com/454270186/CommuTopicPage/repository"
	"github.com/454270186/CommuTopicPage/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewRouter() *gin.Engine {
	dsn := "host=localhost user=postgres password=2021110003 dbname=community port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	ctrl = controller.NewController(service.NewPageService(repository.NewDataRepo(db)))
	router := gin.Default()

	comPage := router.Group("/community/page")
	{
		comPage.GET("/:id", GetTopicById)
		comPage.POST("/topic", CreateNewTopic)
		comPage.POST("/post", CreateNewPostById)
		comPage.DELETE("/topic/:id", DeleteTopicById)
		comPage.DELETE("/post/:id", DeletePostById) // TODO
	}

	return router
}
