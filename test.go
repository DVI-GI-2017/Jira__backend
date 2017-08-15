package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"errors"
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const (
	privKeyPath = "rsa/app.rsa"
	pubKeyPath  = "rsa/app.rsa.pub"
)

var (
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
)

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func initKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

func StartServer() {

	// Non-Protected Endpoint(s)
	http.HandleFunc("/login", LoginHandler)

	http.Handle("/resource", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := ValidateTokenMiddleware(r)

		if err != nil {
			w.WriteHeader(403)
			fmt.Fprintln(w, err)
		} else {
			ProtectedHandler(w, r)
		}
	}))

	fmt.Println("Now listening...")
	http.ListenAndServe(":8080", nil)
}

func StartServerDouble() {

	initKeys()
	StartServer()
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

	response := Response{"Gained access to protected resource"}
	JsonResponse(response, w)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	if strings.ToLower(user.Username) != "someone" {
		if user.Password != "p@ssword" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return
		}
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error extracting the key")
		fatal(err)
	}

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		fatal(err)
	}

	response := Token{tokenString}
	JsonResponse(response, w)

}

func ValidateTokenMiddleware(r *http.Request) (err error) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

	if err == nil {
		if token.Valid {
			return nil
		} else {
			return errors.New("Not valid")
		}
	} else {
		return err
	}
}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
