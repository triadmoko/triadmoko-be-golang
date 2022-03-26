package config

import (
	"fmt"
	"triadmoko-be-golang/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	dsn := "host=ec2-3-225-213-67.compute-1.amazonaws.com user=qjsictgoejtudf password=87913aa93040bdd8730aeba23cac4a983f7ee214cc686fee78e66a217d9f98cb dbname=d1rciihq0fc3n9 port=5432 sslmode=require TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Database not Connect")
	} else {
		fmt.Println("Success connect database")
	}
	db.AutoMigrate(&entity.User{})
	return db
}
