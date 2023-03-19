package activities

import (
	middleware "commerce-platform/middleware"

	"github.com/gin-gonic/gin"
)

func Init(g *gin.RouterGroup) {
	controller := NewController()

	group := g.Group("activities")
	{
		group.GET("", middleware.WithCoreResponseMiddleware(controller.GetHandler))
	}
}
