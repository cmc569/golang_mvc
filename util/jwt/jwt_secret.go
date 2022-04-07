package jwt_secret

import (
	"api/util/crypto"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Token string `json:"token"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(usersId int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	token, _ := crypto.KeyEncrypt(fmt.Sprint(usersId))
	claims := Claims{
		token,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin_extend-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (usersId int, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			usersIdStr, _ := crypto.KeyDecrypt(claims.Token)
			usersId, _ = strconv.Atoi(usersIdStr)
			return usersId, nil
		}
	}

	return 0, err
}
