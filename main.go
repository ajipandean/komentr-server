package main

import (
  "os"
  "fmt"
  "log"
  "github.com/labstack/echo/v4"
  "github.com/joho/godotenv"
  "komentr-server/routes"
)

func main() {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Error on loading .env file")
  }

  e := echo.New()
  routes.Setup(e)

  port := os.Getenv("PORT")
  url := fmt.Sprintf("localhost:%s", port)
  e.Logger.Fatal(e.Start(url))
}
