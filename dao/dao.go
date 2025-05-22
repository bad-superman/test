package dao

import (
	"time"

	"github.com/bad-superman/test/conf"
	"github.com/bad-superman/test/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	c        *conf.Config
	myClient *gorm.DB
}

func New(c *conf.Config) *Dao {
	db, err := gorm.Open(mysql.Open(c.Database.DSN), &gorm.Config{})
	if err != nil {
		logging.Fatalf("db init error: %v", err)
	}
	// 可选：设置连接池参数
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(3000 * time.Second)
	}
	return &Dao{
		c:        c,
		myClient: db,
	}
}
