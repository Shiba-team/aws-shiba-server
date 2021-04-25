package service

import (
	//"bucket/config"
	"bucket/model"
	// "context"
	// "fmt"
	// "os"
	// "time"
	// "github.com/dgrijalva/jwt-go"
	// "go.mongodb.org/mongo-driver/bson"
)

func CreateBucket(input model.Bucket) (bool, string) {
	var name = input.Name
	var username = input.Username
	if !IsExist(input) {
		//Create here
		return true, "Create " + name + " sucessfull by " + username
	} else {
		return false, "Create Faild"
	}
}

func IsExist(input model.Bucket) bool {
	return false
}
