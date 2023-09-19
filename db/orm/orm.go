package orm

import (
	"github.com/tiptophelmet/nomess-core/v5/db/orm/sql"
	"github.com/tiptophelmet/nomess-core/v5/logger"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Init(driver, dsn string) {
	var dialector gorm.Dialector

	switch driver {
	case "mysql":
	case "tidb":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	case "sqlserver":
		dialector = sqlserver.Open(dsn)
	case "clickhouse":
		dialector = clickhouse.Open(dsn)
	default:
		logger.Panic("unsupported db.orm.driver: %v", driver)
	}

	sql.InitGormConnection(dialector, &gorm.Config{TranslateError: true})
}
