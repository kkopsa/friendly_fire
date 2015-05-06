package main

import (
	"testing"
	"gopkg.in/mgo.v2"
   //	"gopkg.in/mgo.v2/bson"
	// "labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
	"log"
	"sync"
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
 
	// Create a wait group to manage the goroutines.
	var waitGroup sync.WaitGroup
 
	// Perform 10 concurrent queries against the database.
	waitGroup.Add(10)
	for query := 0; query < 10; query++ {
		go CreateNewUser("Taylor", "garbagio", &waitGroup, mongoSession)
	}
 
	// Wait for all the queries to complete.
	waitGroup.Wait()
	log.Println("All Queries Completed")
}
