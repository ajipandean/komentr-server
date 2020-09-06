package v1

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "komentr-server/models"
  s "komentr-server/services/api/v1"
)

func RegisterUser(c echo.Context) error {
  user := new(models.User)
  if err := c.Bind(user); err != nil {
    panic(err)
  }
  if err := s.CreateUser(user); err != nil {
    panic(err)
  }
  return c.JSON(http.StatusOK, user)
}
