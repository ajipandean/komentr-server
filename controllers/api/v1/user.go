package controllers_v1

import (
  "fmt"
  "net/http"
  "github.com/labstack/echo/v4"
  "komentr-server/models"
  h "komentr-server/helpers"
  s "komentr-server/services/api/v1"
)

func GetUser(c echo.Context) error {
  u := h.GetLoggedinUser(c)
  return c.JSON(http.StatusOK, u)
}

func GetUserComments(c echo.Context) error {
  user := new(models.User)
  u := h.GetLoggedinUser(c)
  id := uint(u["id"].(float64))
  if err := s.FindUserWithComments(id, user); err != nil {
    fmt.Println(err.Error())
  }
  return c.JSON(http.StatusOK, user)
}
