package activities

import (
	c "commerce-platform/core/config"
	database "commerce-platform/core/database"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Activity struct {
	ID      int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Type    string    `gorm:"column:type;not null"`
	Context string    `gorm:"column:context;not null"`
	Who     string    `gorm:"column:who;not null"`
	When    time.Time `gorm:"column:when;not null"`
}

type IActivityService interface {
	Add(activity Activity) error
	Find(skip uint, take uint) ([]Activity, error)
	WriteActivity(c *gin.Context, activityType string, context any) error
}

type ActivityService struct {
	DBConnectionManager database.IDBConnectionManager
}

func (p *ActivityService) Add(activity Activity) error {
	db, err := p.DBConnectionManager.GetDB()
	if err != nil {
		return err
	}

	err = db.Save(&activity).Error
	return err
}

func (p *ActivityService) Find(skip uint, take uint) ([]Activity, error) {
	db, err := p.DBConnectionManager.GetDB()
	if err != nil {
		return []Activity{}, err
	}

	data := []Activity{}
	err = db.Limit(int(take)).Offset(int(skip)).Find(&data).Error
	return data, err
}

func (p *ActivityService) WriteActivity(c *gin.Context, activityType string, context any) error {
	activityContext, err := json.Marshal(context)
	who, whoError := c.Get("person")
	if err == nil && whoError {
		return p.Add(Activity{
			Type:    activityType,
			Context: string(activityContext),
			Who:     fmt.Sprintf("%v", who),
			When:    time.Now(),
		})
	}
	return nil
}

func NewService() IActivityService {
	config := c.GetConfig()
	manager := database.GetInstance(config)
	return &ActivityService{DBConnectionManager: manager}
}
