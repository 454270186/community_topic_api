package main

import (
	"log"
	"os"

	"github.com/454270186/CommuTopicPage/controller"
	"github.com/454270186/CommuTopicPage/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := repository.InitIndexMap("./data/"); err != nil {
		log.Println(err)
		os.Exit(-1);
	}

	router := gin.Default()

	router.GET("/community/page/:id", func(ctx *gin.Context) {
		topicId := ctx.Param("id")
		data := controller.QueryPageInfo(topicId)
		ctx.JSON(200, data)
	})

	router.Run()
}