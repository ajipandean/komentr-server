package routes

import (
  "github.com/labstack/echo/v4"
  c "komentr-server/controllers/api/v1"
)

func Setup(e *echo.Echo) {
  v1 := e.Group("/api/v1")
  v1.GET("/secret", c.GetAppSecret)
  v1.POST("/auth/register", c.Register)
}
