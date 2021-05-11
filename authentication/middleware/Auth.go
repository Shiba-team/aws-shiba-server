package middleware

import (
	"authentication/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication(Role interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			log.Println("unAuth")
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Token is required!",
			})
			c.Abort()
			return
		}
		token := strings.Split(tokenHeader, " ")[1]
		
		_, err := service.ValidateToken(token);

		if err != nil {
			log.Println("unAuth")
			c.JSON(401, gin.H{
				"error": "No Authentication",
			})
			c.Abort()
			return
		}

		log.Println("Auth")
		c.Next()

	}
}
