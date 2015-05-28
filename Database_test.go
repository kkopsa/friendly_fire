package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"testing"
	"time"
)

func TestCreateNewUser(t *testing.T) {
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Timeout:  60 * time.Second,
		Database: "FriendlyFire",
		//Username: AuthUserName,
		//Password: AuthPassword,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	mongoSession.SetMode(mgo.Monotonic, true)

	// attempt to create a user
	CreateNewUser("Taylor", "garbagio")

	// Wait for all the queries to complete.
	log.Println("All Queries Completed")
}
