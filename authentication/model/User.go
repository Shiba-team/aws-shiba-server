package model

import (
	"authentication/common"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID         primitive.ObjectID `bson:"_id"`
    Name       string            `json:"name" validate:"required,min=2,max=100"`
    Username   string           `json:"username" validate:"required"`
    Email   string           `json:"email" validate:"required"`
    Password   string            `json:"password" validate:"required"`
    Role   string            `json:"role" validate:"required"`
    Created_at time.Time          `json:"created_at"`
    Updated_at time.Time          `json:"updated_at"`
    SecretKey  string             `json:"secretKey"`
}

type UserDto struct {
    Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
    Role string `json:"Role" binding:"required"`
}

func UserDtoToEntity (dto UserDto) interface{}{
    var user User;
    user.Username = dto.Username;
    user.Email = dto.Email;
    password, err := common.HashPassword(dto.Password)
        if(err != nil){
            return err
        }
    user.Password = string(password)
    user.Role = dto.Role;
    user.SecretKey = uuid.NewString()
    user.Updated_at = time.Now()
    if(user.Created_at == time.Time{}){
        user.Created_at = time.Now()
    }
    return user;
}