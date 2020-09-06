package helpers

import (
  "os"
  "time"
  "github.com/dgrijalva/jwt-go"
)

type ClaimsSchema struct {
  Username string
  Email string
  jwt.StandardClaims
}

func ClaimsJWTToken(username, email string) (string, error) {
  claims := &ClaimsSchema{
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
