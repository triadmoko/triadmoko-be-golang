package config

import (
	"fmt"
	"triadmoko-be-golang/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	dsn := "host=ec2-44-194-92-192.compute-1.amazonaws.com user=yuztftsgqvfhbk password=6ed1e7cd312f1c7af7001842031cdb9256d80a3986c399fe16ff2e32497c03ec dbname=dagvse9fdnkcck port=5432 sslmode=require TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Database not Connect")
	} else {
		fmt.Println("Success connect database")
	}
	db.AutoMigrate(&entity.User{})
	return db
}
