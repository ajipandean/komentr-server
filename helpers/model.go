package helpers

import (
  "time"
  "gorm.io/gorm"
)

type Model struct {
  ID uint `json:"id" gorm:"primary_key"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
