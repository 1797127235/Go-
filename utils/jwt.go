package utils

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("dev-secret") // 先写死，之后可放进配置

type Claims struct {
    UserID   uint   `json:"user_id"`
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func GenerateToken(id uint, username string) (string, error) {
    c := Claims{
        UserID:   id,
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "blog-server",
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
    return token.SignedString(JwtSecret)
}
