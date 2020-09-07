package helpers

import (
  "os"
  "time"
  "github.com/dgrijalva/jwt-go"
)

func ClaimsJWTToken(username, email string, id uint) (string, error) {
  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)
  claims["id"] = id
  claims["username"] = username
  claims["email"] = email
  claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
  encodedToken, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
  if err != nil {
    return "", err
  }
  return encodedToken, nil
}
