package controllers_v1

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "komentr-server/models"
  h "komentr-server/helpers"
  s "komentr-server/services/api/v1"
)

func StoreComment(c echo.Context) error {
  comment := new(models.Comment)
  u := h.GetLoggedinUser(c)
  if err := c.Bind(comment); err != nil {
    panic(err)
  }
  comment.UserID = uint(u["id"].(float64))
  if err := s.CreateComment(comment); err != nil {
    panic(err)
  }
  return c.JSON(http.StatusOK, comment)
}
