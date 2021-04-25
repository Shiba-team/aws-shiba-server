package main

import (
	"bucket/config"
	"bucket/controller"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api")
	{
		client.POST("/buckets", controller.CreateBucket)
		client.POST("/buckets/:bucket", controller.Upload)
		client.GET("/buckets/:bucket/*filename", controller.GetFile)
	}

	return r
}

func main() {
	config.ConnectAws()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := setupRouter()
	router.Run(":" + port)
}
