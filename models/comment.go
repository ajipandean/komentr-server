package models

import (
  "gorm.io/gorm"
)

type Comment struct {
  gorm.Model
  Message string
  UserID uint
}
