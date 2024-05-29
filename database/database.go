package database

import (
    "gin-app/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "root:root@tcp(127.0.0.1:3306)/gin_app?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB.AutoMigrate(&models.Product{})
}
