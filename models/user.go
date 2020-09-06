package models

import "komentr-server/helpers"

type User struct {
  helpers.Model
  Username string `json:"username" gorm:"type:varchar(70)"`
  Email string `json:"email" gorm:"type:varchar(100);not null;unique"`
  Password string `json:"password" gorm:"not null;type:varchar(100)"`
  Comments []Comment `json:"comments"`
}
