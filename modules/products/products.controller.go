package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	activities "commerce-platform/modules/activities"
)

type IProductController interface {
	FilterHandler(c *gin.Context) interface{}
	DetailHandler(c *gin.Context) interface{}
}

type ProductController struct {
	Service         IProductService
	Transformer     IProductTransformer
	ActivityService activities.IActivityService
}

func (pc *ProductController) FilterHandler(c *gin.Context) interface{} {
	var query ProductQueryDto
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return nil
	}

	pc.ActivityService.WriteActivity(c, "ProductFilter", query)

	data, err := pc.Service.Filter(query)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil
	}

	return pc.Transformer.TransformMultiple(data)
}

func (pc *ProductController) DetailHandler(c *gin.Context) interface{} {
	idStr, _ := c.Params.Get("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return nil
	}

	pc.ActivityService.WriteActivity(c, "ProductDetail", int(id))

	data, err := pc.Service.FindByID(int(id))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return nil
	}

	return pc.Transformer.Transform(data)
}

func NewController() IProductController {
	return &ProductController{
		Service:         NewService(),
		Transformer:     NewProductTransformer(),
		ActivityService: activities.NewService(),
	}
}
