package main

import (
	"os"

	"authentication/config"
	"authentication/controller"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api")
	{
		client.POST("/auth/register", controller.Register)
	}
	
	return r
}

func main() {
	port := os.Getenv("PORT");
	if port == "" {
		port = "8080"
	}
	router := setupRouter()
	config.ConnectDatabase();
	router.Run(":" + port)
}