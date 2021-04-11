package service

import (
	"authentication/config"
	"authentication/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Register(dto model.UserDto) (i interface{}){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var userCollection *mongo.Collection = config.OpenCollection(config.Client, "user")
	user := model.UserDtoToEntity(dto)
	result, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {	
			fmt.Println("Can not create user!")		
			return nil
		}
		defer cancel()
	return result
}