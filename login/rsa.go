package login

import (
	"fmt"
	"io/ioutil"
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
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
