package modules

import (
	activities "commerce-platform/modules/activities"
	products "commerce-platform/modules/products"

	"github.com/gin-gonic/gin"
)

func InitRouters(g *gin.Engine) {
	apiGroup := g.Group("api")
	{
		products.Init(apiGroup)
		activities.Init(apiGroup)
	}
}
