package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CoreResponseHandlerFunc func(*gin.Context) interface{}

func GetCoreResponseMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Err()
		if err != nil {
			c.JSON(-1, gin.H{
				"result":       -1,
				"errorMessage": err,
			})
		}
	}
}

func WithCoreResponseMiddleware(handler CoreResponseHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := handler(c)
		c.JSON(http.StatusOK, gin.H{
			"result": 0,
			"data":   data,
		})
	}
}
