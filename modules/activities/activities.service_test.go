package activities_test

import (
	"net/http/httptest"
	"regexp"
	"time"

	a "commerce-platform/modules/activities"

	. "commerce-platform/core/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ = Describe("Activities Suite", func() {
	var manager *MockDBConnectionManager
	var activityService a.IActivityService
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		db, m, err := sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())
		mock = m

		dialector := mysql.New(mysql.Config{
			Conn:                      db,
			SkipInitializeWithVersion: true,
		})
		gdb, err := gorm.Open(dialector, &gorm.Config{})
		Expect(err).ShouldNot(HaveOccurred())

		manager = &MockDBConnectionManager{
			gdb: gdb,
		}
		activityService = &a.ActivityService{DBConnectionManager: manager}
		manager.GetDB()
	})

	Describe("Service functions", func() {
		Context("Add an activity", func() {
			It("Should run as expected", func() {
				activity := a.Activity{
					Type:    "Filter",
					Context: "123",
					Who:     "who",
					When:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				}

				sqlInsert := "INSERT INTO `activities` (`type`,`context`,`who`,`when`) VALUES (?,?,?,?)"
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
					WithArgs(activity.Type, activity.Context, activity.Who, activity.When).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				err := activityService.Add(activity)
				Expect(err).To(BeNil())

				err = mock.ExpectationsWereMet()
				Expect(err).To(BeNil())
			})

			It("The WriteActivity function should run as expected", func() {
				context, _ := gin.CreateTestContext(httptest.NewRecorder())
				context.Set("person", "Jerry")

				aContextData := make(map[string]string)
				aContextData["foo"] = "bar"

				sqlInsert := "INSERT INTO `activities` (`type`,`context`,`who`,`when`) VALUES (?,?,?,?)"
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
					WithArgs("Filter", "{\"foo\":\"bar\"}", "Jerry", AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				err := activityService.WriteActivity(context, "Filter", aContextData)
				Expect(err).To(BeNil())

				err = mock.ExpectationsWereMet()
				Expect(err).To(BeNil())
			})
		})

		Context("Find activities", func() {
			It("Should run as expected", func() {
				var skip uint
				var limit uint
				skip, limit = 5, 10

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

				expectedResult := sqlmock.NewRows([]string{"id", "type", "context", "who", "when"})
				for _, activity := range data {
					expectedResult.AddRow(activity.ID, activity.Type, activity.Context, activity.Who, activity.When)
				}

				sqlQuery := "SELECT * FROM `activities` LIMIT 10 OFFSET 5"
				mock.ExpectQuery(regexp.QuoteMeta(sqlQuery)).WillReturnRows(expectedResult)

				result, err := activityService.Find(skip, limit)
				Expect(err).To(BeNil())

				err = mock.ExpectationsWereMet()
				Expect(err).To(BeNil())

				Expect(result).To(Equal(data))
			})
		})
	})
})
