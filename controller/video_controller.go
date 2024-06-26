package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/minhluan1590/Youtube_Golang_Gin_Tutorial/entity"
	"github.com/minhluan1590/Youtube_Golang_Gin_Tutorial/service"
)

// VideoController defines the interface for video-related HTTP operations.
type VideoController interface {
	Save(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

// videoController is a concrete implementation of the VideoController interface.
type videoController struct {
	service service.VideoService
}

// NewVideoController creates a new instance of VideoController.
func NewVideoController(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

// Save handles the HTTP request to save a new video.
func (c *videoController) Save(ctx *gin.Context) {
	var video entity.Video
	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.Save(&video); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, video)
}

// FindAll handles the HTTP request to retrieve all videos.
func (c *videoController) FindAll(ctx *gin.Context) {
	videos, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, videos)
}
