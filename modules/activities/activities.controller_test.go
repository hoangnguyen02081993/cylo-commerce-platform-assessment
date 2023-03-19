package activities_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"

	a "commerce-platform/modules/activities"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Activities Suite", func() {
	var service MockActivitiesService
	var controller a.ActivityController

	BeforeEach(func() {
		service = MockActivitiesService{}
		service.InitMock()
		controller = a.ActivityController{
			Service: &service,
		}

		gin.SetMode(gin.TestMode)
	})

	Describe("Controller", func() {
		Context("GetHandler function", func() {
			It("Should run as expected", func() {
				data := []a.Activity{
					{
						ID:      1,
						Type:    "Filter",
						Context: "123",
						Who:     "who",
						When:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					},
					{
						ID:      2,
						Type:    "Detail",
						Context: "456",
						Who:     "who",
						When:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					},
					{
						ID:      3,
						Type:    "Detail",
						Context: "789",
						Who:     "who",
						When:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					},
				}

				var skip, take uint = 5, 10

				service.SetMockValue("Find", data...)

				context, _ := gin.CreateTestContext(httptest.NewRecorder())
				context.Request = &http.Request{
					Header: make(http.Header),
					URL:    &url.URL{},
				}
				u := url.Values{}
				u.Add("skip", fmt.Sprint(skip))
				u.Add("limit", fmt.Sprint(take))
				context.Request.URL.RawQuery = u.Encode()

				result := controller.GetHandler(context)

				Expect(service.ToBeCalledTime("Find")).To(Equal(1))
				Expect(service.ToBeCalledWith("Find")[0]).To(Equal([]interface{}{
					skip,
					take,
				}))

				Expect(len(context.Errors)).To(Equal(0))
				Expect(result).To(Equal(data))
			})

			It("Not has the query, should run with default value", func() {
				context, _ := gin.CreateTestContext(httptest.NewRecorder())
				context.Request = &http.Request{
					Header: make(http.Header),
					URL:    &url.URL{},
				}

				var skip, take uint = 0, 10

				controller.GetHandler(context)

				Expect(service.ToBeCalledTime("Find")).To(Equal(1))
				calledWith := service.ToBeCalledWith("Find")
				Expect(len(calledWith)).To(Equal(1))

				Expect(calledWith[0]).To(Equal([]interface{}{
					skip, take,
				}))
			})

			It("Should return error when Find return error", func() {
				context, _ := gin.CreateTestContext(httptest.NewRecorder())
				context.Request = &http.Request{
					Header: make(http.Header),
					URL:    &url.URL{},
				}

				resultErr := errors.New("An error")
				service.SetMockError("Find", resultErr)

				var skip, take uint = 0, 10

				result := controller.GetHandler(context)

				Expect(service.ToBeCalledTime("Find")).To(Equal(1))
				Expect(service.ToBeCalledWith("Find")[0]).To(Equal([]interface{}{
					skip,
					take,
				}))

				Expect(result).To(BeNil())
				Expect(len(context.Errors)).To(Equal(1))
				Expect(context.Errors[0].Err).To(Equal(resultErr))
			})
		})
	})
})
