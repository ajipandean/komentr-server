package routes

import (
  "net/http"
  "github.com/labstack/echo/v4"
  c "komentr-server/controllers/api/v1"
)

func Setup(e *echo.Echo) {
  v1 := e.Group("/api/v1")
  v1.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello world")
  })
  v1.POST("/auth/register", c.Register)
}
