package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	jwt.StandardClaims
}

func GenJWTTOken() string {
	mySigningKey := []byte("GiveMeMoney")

	// Create the Claims
	claims := MyCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: (time.Hour * 1).Milliseconds(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return ""
	}
	return ss
}
