package service

import (
	"gin_tut/entity"
)

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

type VideoServiceImpl struct {
	videos []entity.Video
}

func (service *VideoServiceImpl) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *VideoServiceImpl) FindAll() []entity.Video {
	return service.videos
}
