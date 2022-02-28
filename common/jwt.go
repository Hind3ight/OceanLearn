package common

import (
	"errors"
	"github.com/Hind3ight/OceanLearn/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserId uint
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 7

var MySecret = []byte("have2bFun")

func GenToken(user model.User) (string, error) {
	c := &MyClaims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "have2BFun.hind3ight",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(MySecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
