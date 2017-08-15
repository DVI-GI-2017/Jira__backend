package main

import (
	"gopkg.in/mgo.v2"
	//"log"
)

type Context struct {
	Config   *config
	session  *mgo.Session
	database *mgo.Database
}