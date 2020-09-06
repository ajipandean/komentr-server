package main

import (
  "os"
  "fmt"
  "log"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/joho/godotenv"
)

func main() {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Error on loading .env file")
  }

  e := echo.New()

  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello")
  })

  port := os.Getenv("PORT")
  url := fmt.Sprintf("localhost:%s", port)
  e.Logger.Fatal(e.Start(url))
}
