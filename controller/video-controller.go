package controller

import (
	"gin_tut/entity"
	"gin_tut/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type VideoControllerImpl struct {
	Service service.VideoService
}

func (controller *VideoControllerImpl) FindAll() []entity.Video {
	return controller.Service.FindAll()
}

func (controller *VideoControllerImpl) Save(context *gin.Context) entity.Video {
	var video entity.Video
	err := context.BindJSON(&video)
	if err != nil {
		panic(err)
	}
	controller.Service.Save(video)
	return video
}
