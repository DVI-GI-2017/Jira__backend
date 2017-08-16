package tools

import (
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

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
