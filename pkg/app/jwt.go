package app

import (
	"github.com/dgrijalva/jwt-go"
	"golang-admin/global"
	"time"
)

type Claims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userId uint) (interface{}, error) {
	expiredTime := time.Now().Add(global.JWTSetting.Expire)
	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(global.JWTSetting.Secret))
	if err != nil {
		return nil, err
	}

	data := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}
	return data, nil
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.JWTSetting.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}
