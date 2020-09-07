package controllers_v1

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "komentr-server/models"
  s "komentr-server/services/api/v1"
)

func StoreUserComment(c echo.Context) error {
  comment := new(models.Comment)
  if err := c.Bind(comment); err != nil {
    panic(err)
  }
  if err := s.CreateComment(comment); err != nil {
    panic(err)
  }
  return c.JSON(http.StatusOK, comment)
}
