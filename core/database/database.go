package database

import (
	c "commerce-platform/core/config"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IDBConnectionManager interface {
	GetDB() (*gorm.DB, error)
}

type DBConnectionManager struct {
	config     c.AppConfig
	connection *gorm.DB
}

func (m *DBConnectionManager) GetDB() (*gorm.DB, error) {
	if m.connection != nil {
		return m.connection, nil
	}

	db, err := gorm.Open(mysql.Open(m.config.DBConnectionStr), &gorm.Config{})
	if err == nil {
		m.connection = db
	}

	return db, err
}

var singleInstance IDBConnectionManager

func GetInstance(config c.AppConfig) IDBConnectionManager {
	var once sync.Once
	once.Do(func() {
		singleInstance = &DBConnectionManager{
			config: config,
		}
	})
	return singleInstance
}
