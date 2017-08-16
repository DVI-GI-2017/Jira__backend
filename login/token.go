package login

import (
	"fmt"
	"time"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type Token struct {
	Token string `json:"token"`
}

func NewToken() (Token, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := make(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()

	token.Claims = claims

	result, err := token.SignedString(SignKey)

	if err != nil {
		return Token{result}, fmt.Errorf("can not create signed string: %v", err)
	}

	return Token{result}, nil
}

func ValidateToken(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return VerifyKey, nil
			})

		if err == nil {
			if token.Valid {
				h(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Unauthorized")
			}
		}
	}
}