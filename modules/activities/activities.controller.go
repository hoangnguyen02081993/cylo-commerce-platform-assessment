package activities

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IActivityController interface {
	GetHandler(c *gin.Context) interface{}
}

type ActivityController struct {
	Service IActivityService
}

func (a *ActivityController) GetHandler(c *gin.Context) interface{} {
	var query ActivityQueryDto
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return nil
	}

	data, err := a.Service.Find(query.Skip, query.Take)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return nil
	}

	return data
}

func NewController() ActivityController {
	return ActivityController{
		Service: NewService(),
	}
}
