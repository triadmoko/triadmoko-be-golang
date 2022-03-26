package config

import (
	"fmt"
	"os"
	"triadmoko-be-golang/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	dsn := os.Getenv("USER") + ":" + os.Getenv("PASSWORD") + "@tcp(" + os.Getenv("HOST") + ":" + os.Getenv("PORTDB") + ")/" + os.Getenv("DBNAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Database not Connect")
	} else {
		fmt.Println("Success connect database")
	}
	db.AutoMigrate(&entity.User{})
	return db
}
