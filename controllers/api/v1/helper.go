package controllers_v1

import (
  "net/http"
  "github.com/labstack/echo/v4"
  h "komentr-server/helpers"
)

func GetAppSecret(c echo.Context) error {
  return c.String(http.StatusOK, h.GenerateRandomString(32))
}
