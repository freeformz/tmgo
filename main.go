package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Database: "123",
		Username: "456",
		Password: "789",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(session)
}
