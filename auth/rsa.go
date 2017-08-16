package auth

import (
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

func InitKeys() error {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		return err
	}

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return err
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return err
	}

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return err
	}

	return nil
}
