package products_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	p "commerce-platform/modules/products"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product Suite", func() {
	var productService MockProductServiceService
	var activityService MockActivitiesService
	var transformer p.IProductTransformer
	var controller p.IProductController

	BeforeEach(func() {
		productService = MockProductServiceService{}
		productService.InitMock()
		activityService = MockActivitiesService{}
		activityService.InitMock()
		transformer = p.NewProductTransformer()

		controller = &p.ProductController{
			Service:         &productService,
			ActivityService: &activityService,
			Transformer:     transformer,
		}

		gin.SetMode(gin.TestMode)
	})

	Describe("Controller", func() {
		Context("Filter handler", func() {
			It("Should run as expected", func() {
				data := []p.Product{
					{
						ID:          1,
						Name:        "Name 1",
						Description: "Description 1",
						ImgURL:      "ImageUrl 1",
						Price:       14,
						Branch:      "Branch 1",
					},
					{
						ID:          2,
						Name:        "Name 2",
						Description: "Description 2",
						ImgURL:      "ImageUrl 2",
						Price:       30,
						Branch:      "Branch 2",
					},
				}

				productQueryDto := p.ProductQueryDto{
					Filter: p.ProductFilterDto{
						Name:      "bia",
						FromPrice: 0,
						ToPrice:   20,
						Branch:    "Milk",
					},
					Sort: p.ProductSortDto{
						SortField:     "Name",
						SortDirection: "ASC",
					},
				}

				productService.MockTool.SetMockValue("Filter", data...)

				context, _ := gin.CreateTestContext(httptest.NewRecorder())
				context.Request = &http.Request{
					Header: make(http.Header),
					URL:    &url.URL{},
				}
				u := url.Values{}
				u.Add("name", fmt.Sprint(productQueryDto.Filter.Name))
				u.Add("fromPrice", fmt.Sprint(productQueryDto.Filter.FromPrice))
				u.Add("toPrice", fmt.Sprint(productQueryDto.Filter.ToPrice))
				u.Add("branch", fmt.Sprint(productQueryDto.Filter.Branch))
				u.Add("sortField", fmt.Sprint(productQueryDto.Sort.SortField))
				u.Add("sortDirection", fmt.Sprint(productQueryDto.Sort.SortDirection))
				context.Request.URL.RawQuery = u.Encode()

				result := controller.FilterHandler(context)

				Expect(productService.ToBeCalledTime("Filter")).To(Equal(1))
				Expect(productService.ToBeCalledWith("Filter")[0]).To(Equal([]interface{}{
					productQueryDto,
				}))

				Expect(activityService.ToBeCalledTime("WriteActivity")).To(Equal(1))
				Expect(activityService.ToBeCalledWith("WriteActivity")[0]).To(Equal([]interface{}{
					context,
					"ProductFilter",
					productQueryDto,
				}))

				Expect(len(context.Errors)).To(Equal(0))
				Expect(result).To(Equal(transformer.TransformMultiple(data)))
			})

			It("Without query, should run as expected", func() {
				context, _ := gin.CreateTestContext(httptest.NewRecorder())
				context.Request = &http.Request{
					Header: make(http.Header),
					URL:    &url.URL{},
				}

				productQueryDto := p.ProductQueryDto{
					Filter: p.ProductFilterDto{
						Name:      "",
						FromPrice: 0,
						ToPrice:   0,
						Branch:    "",
					},
					Sort: p.ProductSortDto{
						SortField:     "",
						SortDirection: "",
					},
				}

				controller.FilterHandler(context)

				Expect(productService.ToBeCalledTime("Filter")).To(Equal(1))
				Expect(productService.ToBeCalledWith("Filter")[0]).To(Equal([]interface{}{
					productQueryDto,
				}))

				Expect(len(context.Errors)).To(Equal(0))
			})

			It("Should throw error when filter data error", func() {
				context, _ := gin.CreateTestContext(httptest.NewRecorder())
				context.Request = &http.Request{
					Header: make(http.Header),
					URL:    &url.URL{},
				}

				err := errors.New("An error")
				productService.SetMockError("Filter", err)

				controller.FilterHandler(context)

				Expect(len(context.Errors)).To(Equal(1))
				Expect(context.Errors[0].Err).To(Equal(err))
			})
		})

		Context("DetailHandler", func() {
			It("Should run as expected", func() {
				data := p.Product{
					ID:          1,
					Name:        "Name 1",
					Description: "Description 1",
					ImgURL:      "ImageUrl 1",
					Price:       14,
					Branch:      "Branch 1",
				}

				productService.MockTool.SetMockValue("FindByID", data)

				context, _ := gin.CreateTestContext(httptest.NewRecorder())
				context.Request = &http.Request{
					Header: make(http.Header),
					URL:    &url.URL{},
				}

				var id int = 1
				context.Params = []gin.Param{
					{
						Key:   "id",
						Value: fmt.Sprint(id),
					},
				}
				result := controller.DetailHandler(context)

				Expect(len(context.Errors)).To(Equal(0))

				Expect(activityService.ToBeCalledTime("WriteActivity")).To(Equal(1))
				Expect(activityService.ToBeCalledWith("WriteActivity")[0]).To(Equal([]interface{}{
					context,
					"ProductDetail",
					id,
				}))

				Expect(productService.ToBeCalledTime("FindByID")).To(Equal(1))
				Expect(productService.ToBeCalledWith("FindByID")[0]).To(Equal([]interface{}{
					id,
				}))

				Expect(result).To(Equal(transformer.Transform(data)))
			})

			It("Should throw error when find data error", func() {
				context, _ := gin.CreateTestContext(httptest.NewRecorder())
				context.Request = &http.Request{
					Header: make(http.Header),
					URL:    &url.URL{},
				}

				var id int = 1
				context.Params = []gin.Param{
					{
						Key:   "id",
						Value: fmt.Sprint(id),
					},
				}

				err := errors.New("An error")
				productService.SetMockError("FindByID", err)

				controller.DetailHandler(context)

				Expect(len(context.Errors)).To(Equal(1))
				Expect(context.Errors[0].Err).To(Equal(err))
			})
		})
	})
})
