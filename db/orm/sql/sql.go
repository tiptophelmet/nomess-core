package sql

import (
	"sync"

	"github.com/tiptophelmet/nomess-core/v3/logger"
	"gorm.io/gorm"
)

type gormClient struct {
	conn *gorm.DB
	mu   sync.Mutex
}

var client *gormClient

func InitGormConnection(dialector gorm.Dialector, config *gorm.Config) {
	gormConn, err := gorm.Open(dialector, config)
	if err != nil {
		logger.Panic("failed to resolve connect to postgres db: %v", err.Error())
	}

	client = &gormClient{conn: gormConn}
	prepareConnectionPool()
}

func prepareConnectionPool() {
	sqlDB, err := client.conn.DB()
	if err != nil {
		logger.Panic("failed to resolve connect to db: %v", err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func Connection() *gorm.DB {
	client.mu.Lock()
	defer client.mu.Unlock()

	return client.conn
}
