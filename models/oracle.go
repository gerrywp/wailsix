package models

import (
	"fmt"
	"time"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	// 构造 DSN
	url := oracle.BuildUrl(
		"172.16.34.7",
		1521,
		"testdb",
		"mestest",
		"mestest",
		nil,
	)

	// 直接用 oracle.Open（这个是存在的）
	dialector := oracle.Open(url)

	var err error
	DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("gorm open failed: %w", err)
	}

	// 连接池
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return nil
}