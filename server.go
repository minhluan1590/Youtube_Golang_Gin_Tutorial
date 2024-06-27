package main

import (
	"github.com/gin-gonic/gin"
	"github.com/minhluan1590/Youtube_Golang_Gin_Tutorial/controller"
	"github.com/minhluan1590/Youtube_Golang_Gin_Tutorial/middlewares"
	"github.com/minhluan1590/Youtube_Golang_Gin_Tutorial/service"
	gindump "github.com/tpkeeper/gin-dump"
	"net/http"
)

func main() {
	r := gin.New()

	// Add middlewares
	r.Use(middlewares.Logger(), gin.Recovery(), middlewares.BasicAuth(), gindump.Dump())

	// Initialize services and controllers
	videoService := service.NewVideoService()
	videoController := controller.NewVideoController(videoService)

	// Define routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// API group
	api := r.Group("/api")
	{
		api.POST("/videos", videoController.Save)
		api.GET("/videos", videoController.FindAll)
	}

	// Serve the index.html file
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		videos, err := videoService.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":  "Home Page",
			"videos": videos,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
