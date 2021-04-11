package controller

import (
	"authentication/model"
	"authentication/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)



func Register(c * gin.Context){
	var input model.UserDto;
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	  }
	var result interface{}= service.Register(input);
	if(result == nil){
		msg := fmt.Sprintf("Can not create user!")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}