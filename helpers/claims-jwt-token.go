package helpers

import (
  "os"
  "time"
  "github.com/dgrijalva/jwt-go"
)

type ClaimsSchema struct {
  ID uint `json:"id"`
  Username string `json:"username"`
  Email string `json:"email"`
  jwt.StandardClaims
}

func ClaimsJWTToken(username, email string, id uint) (string, error) {
  claims := &ClaimsSchema{
    id,
    username,
    email,
    jwt.StandardClaims{
      IssuedAt: time.Now().Unix(),
    },
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  appSecret := os.Getenv("APP_SECRET")
  encodedToken, err := token.SignedString([]byte(appSecret))
  if err != nil {
    return "", err
  }
  return encodedToken, nil
}
