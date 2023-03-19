package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthenticationMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		person := ctx.Request.Header.Get("X-Person")
		if person == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"result":       -1,
				"errorMessage": "You need to login first",
			})
			return
		}

		ctx.Set("person", person)
		ctx.Next()
	}
}
