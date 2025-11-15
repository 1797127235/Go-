package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
   jwt
   jwt是基于token的认证机制
*/

var JwtSecret = []byte("dev-secret") // 先写死，之后可放进配置

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成token
func GenerateToken(id uint, username string) (string, error) {
	c := Claims{
		UserID:   id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{ //生成token的注册信息
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
			Issuer:    "blog-server",
		},
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(JwtSecret)
}

// 解析并验证 token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})


	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
