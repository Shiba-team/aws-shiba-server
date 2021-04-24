package model

// import (
// 	"time"
// 	//"file/model"
	
// 	"github.com/google/uuid"
// )

type Bucket struct {
    Username   string           `json:"username" validate:"required"`
	Name   string            `json:"name" validate:"required"`
	
    //List_file
    //Role   string            `json:"role" validate:"required"`
//
    //SecretKey  string             `json:"secretKey"`
}

