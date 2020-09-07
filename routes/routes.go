package routes

import (
  "os"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  h "komentr-server/helpers"
  c "komentr-server/controllers/api/v1"
)

func Setup(e *echo.Echo) {
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"http://localhost:3000"},
  }))

  v1 := e.Group("/v1")
  v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
    Format: "method=${method}, uri=${uri}, status=${status}\n",
  }))
  v1.Use(middleware.Recover())
  v1.GET("/secret", c.GetAppSecret)
  v1.POST("/auth/login", c.Login)
  v1.POST("/auth/register", c.Register)

  v1s := v1.Group("/secure")
  jwtConfig := middleware.JWTConfig{
    Claims: &h.ClaimsSchema{},
    SigningKey: []byte(os.Getenv("APP_SECRET")),
  }
  v1s.Use(middleware.JWTWithConfig(jwtConfig))
  v1s.GET("", c.GetUser)
  v1s.POST("/comments", c.StoreComment)
}
