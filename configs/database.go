package configs

import (
	"database/sql"
	"fmt"
	"gohub/initPack/config"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认数据库mysql
			"connection": config.Env("DB_CONNECTION", "mysql"),

			// MySQL连接信息
			"mysql": map[string]interface{}{
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", "3306"),
				"database": config.Env("DB_DATABASE", "gohub"),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"charset":  "utf8mb4",

				// 连接池配置
				"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
				"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
				"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 5*60),
			},
		}
	})
}

// DB 对象
var DB *gorm.DB
var SQLDB *sql.DB

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, _logger gormLogger.Interface) {

	// 使用 gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

}
