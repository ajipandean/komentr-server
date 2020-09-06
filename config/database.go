package config

import (
  "os"
  "fmt"
  "log"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "komentr-server/models"
)

var DB *gorm.DB

func ConnectToDatabase() {
  var err error
  dataSource := fmt.Sprintf(
    "%s:%s@%s/%s?charset=utf8mb4&parseTime=true&loc=Local",
    os.Getenv("DATABASE_USERNAME"),
    os.Getenv("DATABASE_PASSWORD"),
    os.Getenv("DATABASE_HOST"),
    os.Getenv("DATABASE_NAME"),
  )
  DB, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})
  if err != nil {
    log.Fatal("Failed to connect to database")
  } else {
    DB.AutoMigrate(&models.User{}, &models.Comment{})
    fmt.Println("Connected to database")
  }
}
