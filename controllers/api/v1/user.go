package controllers_v1

import (
  "net/http"
  "github.com/labstack/echo/v4"
  h "komentr-server/helpers"
)

func GetUser(c echo.Context) error {
  u := h.GetLoggedinUser(c)
  return c.JSON(http.StatusOK, u)
}
