package config

import (
	"DDD/app/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=yuleyek dbname=travellingApps port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database Connected")
	db.AutoMigrate(&entity.User{}, &entity.Permission{}, &entity.Role{})
	return db
}
