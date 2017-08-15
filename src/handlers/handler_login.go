package handlers

import (
"net/http"
"encoding/json"
"request"
"fmt"
"common/response"
"gopkg.in/mgo.v2/bson"
"crypto/sha1"
"common/model"
"router"
"log"
)

func passwHash(password string) string {
	const authSalt = "DUTLtcDJppHi8ZPBBxypEA7o1yN0pMbv"

	return fmt.Sprintf("%x", sha1.Sum([]byte(password+authSalt)))
}

func LoginRequest(r *http.Request) (out router.Response) {
	var (
		rq   request.LoginRequest
		user model.Client
	)

	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		log.Printf("Failed to decode LoginRequest body from host %s:\n%s", r.RemoteAddr, r.Body)
		out.Error = fmt.Errorf("Couldn't decode request: %s", err.Error())
		return
	}

	if err := ctx.Database().C("clients").Find(bson.M{
		"user":     rq.User,
		"password": passwHash(rq.Password),
	}).One(&user); err == nil {
		out.Body = response.NewLoginResponse(user.Token)
	} else {
		log.Printf("Auth failed for user '%s' from host %s: %s", rq.User, r.RemoteAddr, err)
		out.Error = fmt.Errorf("Пользователь с указанным именем и паролем не найден")
	}

	return
}