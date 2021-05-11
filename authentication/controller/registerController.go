package controller

import (
	"authentication/model"
	"authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)



func Register(c * gin.Context){
	var input model.UserRegisterDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	  }
	err := service.Register(input)
	if err != ""{
		c.JSON(http.StatusInternalServerError, gin.H{"messageError": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Register successful!"})
}