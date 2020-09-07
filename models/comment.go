package models

import (
  "komentr-server/helpers"
)

type Comment struct {
  helpers.Model
  Message string `json:"message"`
  UserID uint `json:"user_id"`
}
