package models

import (
  "strings"
  "komentr-server/helpers"
  "gorm.io/gorm"
  b "golang.org/x/crypto/bcrypt"
)

type User struct {
  helpers.Model
  Username string `json:"username" gorm:"type:varchar(70);"`
  Email string `json:"email" gorm:"type:varchar(100);not null;unique"`
  Password string `json:"password" gorm:"not null;type:varchar(100)"`
  Comments []Comment `json:"comments"`
}

func (u *User) ValidatePassword(password string) error {
  hashedPassword := []byte(u.Password)
  currentPassword := []byte(password)
  err := b.CompareHashAndPassword(hashedPassword, currentPassword)
  if err != nil {
    return err
  }
  return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
  bytePassword := []byte(u.Password)
  hashedPassword, err := b.GenerateFromPassword(bytePassword, b.DefaultCost)
  if err != nil {
    panic(err)
  }
  u.Password = string(hashedPassword)

  splittedEmail := strings.Split(u.Email, "@")
  u.Username = splittedEmail[0]
  return nil
}
