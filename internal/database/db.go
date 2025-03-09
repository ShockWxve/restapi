package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=pwd1111! dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	// Убираем миграцию для User, чтобы не было конфликтов с ограничениями
	// DB.AutoMigrate(&taskService.Task{})
	// Если нужно добавить Task и User вместе, раскомментируй:
	// DB.AutoMigrate(&taskService.Task{}, &userService.User{})
}
