package common

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Openid string
	jwt.StandardClaims
}

var hmacSampleSecret []byte

// token秘钥初始化
func JwtInit(hmacSecret []byte) {
	hmacSampleSecret = hmacSecret
}

// 解析创建
func CreateToken(openid string) (string, error) {

	// 设置令牌到期时间为5分钟
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{Openid: openid, StandardClaims: jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}

// 解析token
func ParseToken(tokenString string) (string, error) {

	// Parse 放回 Token
	// 方法是用来校验加密 和过期时间的
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	claims, ok := token.Claims.(Claims)

	if ok && token.Valid {
		return claims.Openid, nil
	}

	return "", err
}
