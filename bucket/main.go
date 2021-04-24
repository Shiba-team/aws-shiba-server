package main

import (
	"os"
	"bucket/controller"
	"github.com/gin-gonic/gin"
	
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api")
	{
		client.POST("/bucket/create", controller.CreateBucket)
	}
	
	return r
}

func main() {
	port := os.Getenv("PORT");
	if port == "" {
		port = "8080"
	}
	router := setupRouter()
	router.Run(":" + port)
}