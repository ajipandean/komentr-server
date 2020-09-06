package service_v1

import (
  "komentr-server/models"
  "komentr-server/config"
)

func CreateUser(user *models.User) error {
  db := config.DB
  result := db.Create(&user)
  return result.Error
}
