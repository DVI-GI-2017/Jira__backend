package tools

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken() (result string, err error) {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	result, err = token.SignedString(SignKey)

	if err != nil {
		return
	}

	return result, nil
}
