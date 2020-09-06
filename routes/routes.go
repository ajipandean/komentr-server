package routes

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
  v1 := e.Group("/api/v1")
  v1.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello world")
  })
}
