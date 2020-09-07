package routes

import (
  "os"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  c "komentr-server/controllers/api/v1"
)

func Setup(e *echo.Echo) {
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"http://localhost:3000"},
  }))
  e.Use(middleware.Recover())
  e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
    Format: "method=${method}, uri=${uri}, status=${status}\n",
  }))

  v1 := e.Group("/v1")
  v1.GET("/secret", c.GetAppSecret)
  v1.POST("/auth/login", c.Login)
  v1.POST("/auth/register", c.Register)

  s := v1.Group("/secure")
  s.Use(middleware.JWT([]byte(os.Getenv("APP_SECRET"))))
  s.GET("/user", c.GetUser)
  s.GET("/user/comments", c.GetUserComments)
  s.POST("/comments", c.StoreComment)
}
