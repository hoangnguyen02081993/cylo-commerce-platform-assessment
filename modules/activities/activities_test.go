package activities_test

import (
	"testing"

	a "commerce-platform/modules/activities"
	m "commerce-platform/test"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

func TestActivity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Activities Suite")
}

type MockDBConnectionManager struct {
	gdb *gorm.DB
}

func (m *MockDBConnectionManager) GetDB() (*gorm.DB, error) {
	return m.gdb, nil
}

type MockActivitiesService struct {
	m.MockTool[a.Activity]
}

func (p *MockActivitiesService) InitMock() {
	p.MockTool = m.MockTool[a.Activity]{}
	p.MockTool.InitMock()
}

func (p *MockActivitiesService) Add(activity a.Activity) error {
	funcName := "Add"
	p.MockTool.SetMockFunction(funcName, activity)
	return p.MockTool.GetMockError(funcName)
}

func (p *MockActivitiesService) Find(skip uint, take uint) ([]a.Activity, error) {
	funcName := "Find"
	p.MockTool.SetMockFunction(funcName, skip, take)
	return p.MockTool.GetMockValue(funcName), p.MockTool.GetMockError(funcName)
}

func (p *MockActivitiesService) WriteActivity(c *gin.Context, activityType string, context any) error {
	funcName := "WriteActivity"
	p.MockTool.SetMockFunction(funcName, c, activityType, context)
	return p.MockTool.GetMockError(funcName)
}
