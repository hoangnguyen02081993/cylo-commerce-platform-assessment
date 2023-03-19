package products

import (
	"commerce-platform/middleware"

	"github.com/gin-gonic/gin"
)

func Init(g *gin.RouterGroup) {
	controller := NewController()

	group := g.Group("products")
	{
		group.GET("", middleware.WithCoreResponseMiddleware(controller.FilterHandler))
		group.GET(":id", middleware.WithCoreResponseMiddleware(controller.DetailHandler))
	}
}
