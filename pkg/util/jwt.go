package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtUserSecret []byte
var jwtAdminSecret []byte

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

// GenerateUserToken generate tokens used for auth
func GenerateUserToken(id int) (string, error) {
	return generateToken(id,jwtUserSecret)
}
func GenerateAdminToken(id int) (string, error) {
	return generateToken(id,jwtAdminSecret)
}

// ParseUserToken parsing token
func ParseUserToken(token string) (*Claims, error) {
	return parseToken(token,jwtUserSecret)
}

// ParseUserToken parsing token
func ParseAdminToken(token string) (*Claims, error) {
	return parseToken(token,jwtAdminSecret)
}

func generateToken(id int,secret []byte)  (string, error)  {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 30 * time.Hour)

	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "lifang.biz",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secret)

	return token, err
}


// ParseUserToken parsing token
func parseToken(token string,secret []byte) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
