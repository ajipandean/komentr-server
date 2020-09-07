package controllers_v1

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/dgrijalva/jwt-go"
  h "komentr-server/helpers"
)

func GetUser(c echo.Context) error {
  user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*h.ClaimsSchema)
  return c.JSON(http.StatusOK, claims)
}
