package service_v1

import (
  "komentr-server/models"
  "komentr-server/config"
)

func CreateUser(user *models.User) error {
  db := config.DB
  result := db.Create(&user)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func VerifyUser(username, password string, user *models.User) error {
  db := config.DB
  result := db.Where("username = ?", username).First(&user)
  if result.Error != nil {
    return result.Error
  }
  if err := user.ValidatePassword(password); err != nil {
    return err
  }
  return nil
}
