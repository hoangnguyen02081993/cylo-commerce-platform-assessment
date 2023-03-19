package products_test

import (
	"testing"

	a "commerce-platform/modules/activities"
	p "commerce-platform/modules/products"
	m "commerce-platform/test"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

func TestProduct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Product Suite")
}

type MockDBConnectionManager struct {
	gdb  *gorm.DB
	mock sqlmock.Sqlmock
}

func (m *MockDBConnectionManager) GetDB() (*gorm.DB, error) {
	return m.gdb, nil
}

type MockProductServiceService struct {
	m.MockTool[p.Product]
}

func (s *MockProductServiceService) InitMock() {
	s.MockTool = m.MockTool[p.Product]{}
	s.MockTool.InitMock()
}

func (s *MockProductServiceService) Filter(query p.ProductQueryDto) ([]p.Product, error) {
	funcName := "Filter"
	s.MockTool.SetMockFunction(funcName, query)
	return s.MockTool.GetMockValue(funcName), s.MockTool.GetMockError(funcName)
}

func (s *MockProductServiceService) FindByID(id int) (p.Product, error) {
	funcName := "FindByID"
	s.MockTool.SetMockFunction(funcName, id)
	value := s.MockTool.GetMockValue(funcName)
	if value == nil {
		return p.Product{}, s.MockTool.GetMockError(funcName)
	}
	return value[0], s.MockTool.GetMockError(funcName)
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
