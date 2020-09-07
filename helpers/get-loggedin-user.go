package helpers

import (
  "github.com/labstack/echo/v4"
  "github.com/dgrijalva/jwt-go"
)

func GetLoggedinUser(c echo.Context) jwt.Claims {
  user := c.Get("user").(*jwt.Token)
  claims := user.Claims
  return claims
}
