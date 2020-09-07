package controllers_v1

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "komentr-server/models"
  h "komentr-server/helpers"
  s "komentr-server/services/api/v1"
)

type loginSchema struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

func Register(c echo.Context) error {
  user := new(models.User)
  if err := c.Bind(user); err != nil {
    panic(err)
  }
  if err := s.CreateUser(user); err != nil {
    panic(err)
  }
  return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
  user := new(models.User)
  schema := new(loginSchema)
  if err := c.Bind(schema); err != nil {
    panic(err)
  }
  if err := s.VerifyUser(schema.Username, schema.Password, user); err != nil {
    panic(err)
  }
  token, err := h.ClaimsJWTToken(user.Username, user.Email, user.ID)
  if err != nil {
    panic(err)
  }
  return c.JSON(http.StatusOK, token)
}
