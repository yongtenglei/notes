package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

//var
//TokenExpired     = errors.New("Token 已过期")
//TokenNotValidYet = errors.New("Token 不再有效")
//TokenMalFormed   = errors.New("Token 非法")
//TokenInvalid     = errors.New("Token 无效")
//)

const (
	secretKey   = "secretkey"
	ExpiresTime = 7 * 24 * time.Hour
)

type CustomClaims struct {
	jwt.StandardClaims
	ID       uint64
	UserName string
}

type JWT struct {
	SignedKey []byte
}

func NewJWT() *JWT {
	return &JWT{SignedKey: []byte(secretKey)}
}

// 创建Token
func (j *JWT) GenerateToken(claims CustomClaims) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(j.SignedKey)
	return token, err
}

// 解析Token
func (j *JWT) ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignedKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 刷新Token
func (j *JWT) RefreshToken(token string) (string, error) {

	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignedKey, nil
	})

	if err != nil {
		return "", err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			claims.StandardClaims.ExpiresAt = time.Now().Add(ExpiresTime).Unix()
			return j.GenerateToken(*claims)
		}
	}

	return "", errors.New("Invalid Token")

}
