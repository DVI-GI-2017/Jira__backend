package handlers

import (
"net/http"
"encoding/json"
"common/request"
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

