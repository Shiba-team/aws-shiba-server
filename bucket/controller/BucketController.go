package controller

import (
	"net/http"
	"bucket/service"
	"bucket/model"
	"github.com/gin-gonic/gin"
)

func CreateBucket(c * gin.Context){
	var input model.Bucket;
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resutl, msg := service.CreateBucket(input)
	if(resutl){
		c.JSON(http.StatusOK, gin.H{"Result": msg})
		return
	}else{
		c.JSON(http.StatusBadRequest, gin.H{"Result" :msg})
		return
	}
}