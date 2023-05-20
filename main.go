package main

import (
	"github.com/454270186/CommuTopicPage/controller"
	"github.com/454270186/CommuTopicPage/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := repository.InitIndexMap("./data/"); err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/community/page/:id", func(ctx *gin.Context) {
		topicId := ctx.Param("id")
		data := controller.QueryPageInfo(topicId)
		ctx.JSON(200, data)
	})

	router.Run()
}