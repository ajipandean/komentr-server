package service_v1

import (
  "komentr-server/models"
  "komentr-server/config"
)

func FetchUserComments(id uint, comments *[]models.Comment) error {
  db := config.DB
  result := db.Where("user_id = ?", id).Find(&comments)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func CreateComment(comment *models.Comment) error {
  db := config.DB
  result := db.Create(&comment)
  if result.Error != nil {
    return result.Error
  }
  return nil
}
