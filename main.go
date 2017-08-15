//package main
//
//import (
//	"net/http"
//	"fmt"
//	"strconv"
//	"log"
//	"Jira__backend/tools"
//	"Jira__backend/router"
//	"Jira__backend/dataBase"
//)
//
//func main() {
//	dataBase.NewDBConnection()
//
//	port, err := tools.GetServerPort("configs/server.json")
//
//	if err != nil {
//		panic("bad config")
//	}
//
//	router.NewRouter()
//
//	fmt.Printf("Server started on port %d...\n", port)
//	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
//}
package main

//  Generate RSA signing files via shell (adjust as needed):
//
//  $ openssl genrsa -out app.rsa 1024
//  $ openssl rsa -in app.rsa -pubout > app.rsa.pub
//
// Code borrowed and modified from the following sources:
// https://www.youtube.com/watch?v=dgJFeqeXVKw
// https://goo.gl/ofVjK4
// https://github.com/dgrijalva/jwt-go
//

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"errors"
)

const (
	// For simplicity these files are in the same folder as the app binary.
	// You shouldn't do this in production.
	privKeyPath = "app.rsa"
	pubKeyPath  = "app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
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

	// Protected Endpoints
	http.Handle("/resource", http.HandlerFunc(ProtectedHandler))

	log.Println("Now listening...")
	http.ListenAndServe(":8080", nil)
}

func main() {

	initKeys()
	StartServer()
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

	err := ValidateTokenMiddleware(w, r)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)

		return
	}

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

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request) (err error) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

	if err == nil {
		if token.Valid {
			return nil
		}

		err = errors.New("Not valid")
		return
	}

	return nil
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
