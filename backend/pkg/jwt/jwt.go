package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)


type Claims struct {
  jwt.StandardClaims
  UserId uint `json:"user_id"`
  IsAdmin bool `json:"is_admin"`
}

func GenerateToken(userId uint, isAdmin bool) (string, error) {
  claims := Claims{
    StandardClaims: jwt.StandardClaims{
      Issuer: fmt.Sprintf("%s/backend", os.Getenv("PUBLIC_APP_NAME")),
      ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
    },
    UserId: userId,
    IsAdmin: isAdmin,
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNATURE_SECRET")))
  if err != nil {
    return "", err
  }

  return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
  return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS256 {
      return nil, fmt.Errorf("there was an error")
    }
    return []byte(os.Getenv("JWT_SIGNATURE_SECRET")), nil
  })
}

func GetUserIdFromToken(token *jwt.Token) (uint, bool, error) {
  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    userId := claims["user_id"].(float64)
    isAdmin := claims["is_admin"].(bool)
    return uint(userId), isAdmin, nil
  }
  return 0, false, fmt.Errorf("invalid token")
}
