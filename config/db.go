package config

import (
	"fmt"
	"triadmoko-be-golang/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	dsn := "host=ec2-3-225-213-67.compute-1.amazonaws.com user=caivnfhfypngwn password=9a743a329c1b0855cb8504b551432c18861974de50ab9115102aa22b635d1226 dbname=dahc9ia5r6dt0m port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Database not Connect")
	} else {
		fmt.Println("Success connect database")
	}
	db.AutoMigrate(&entity.User{})
	return db
}
