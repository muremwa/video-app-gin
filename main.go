package main

import (
	"gin_tut/controller"
	"gin_tut/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    = service.VideoServiceImpl{}
	videoController = controller.VideoControllerImpl{Service: &videoService}
)

func main() {
	r := gin.Default()

	r.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})

	r.POST("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})

	err := r.Run(":8000")
	if err != nil {
		return
	}
}
