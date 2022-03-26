package config

import (
	"fmt"
	"os"
	"triadmoko-be-golang/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	dsn := "host=" + os.Getenv("HOST") + " user=" + os.Getenv("USER") + " password=" + os.Getenv("PASSWORD") + " dbname=" + os.Getenv("DBNAME") + " port=" + os.Getenv("PORTDB") + " sslmode=enable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Database not Connect")
	} else {
		fmt.Println("Success connect database")
	}
	db.AutoMigrate(&entity.User{})
	return db
}
