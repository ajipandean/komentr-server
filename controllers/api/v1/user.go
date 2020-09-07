package controllers_v1

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "komentr-server/models"
  h "komentr-server/helpers"
  s "komentr-server/services/api/v1"
)

func GetUser(c echo.Context) error {
  u := h.GetLoggedinUser(c)
  return c.JSON(http.StatusOK, u["username"])
}

func GetUserComments(c echo.Context) error {
  user := new(models.User)
  if err := s.FindUserWithComments(2, user); err != nil {
    panic(err)
  }
  return c.JSON(http.StatusOK, user)
}
