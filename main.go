package main

import (
	"log"
	"net/http"
	"time"
	"gopkg.in/mgo.v2"
//	"sync"
)

var MongoSession *mgo.Session
//var WaitGroup sync.WaitGroup

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		inner.ServeHTTP(w, r)
		
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// Create a connection to the database
func startDbSession() {
	// db connection information
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Timeout:  60 * time.Second,
		Database: "FriendlyFire",
		//Username: AuthUserName,
		//Password: AuthPassword,
	}
	
	// connect
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n")
		panic("Can't connect to database\n")
	}
	mongoSession.SetMode(mgo.Monotonic, true)
	
	// TODO probably a bad design to have the database session globally available
	// so change this whole thing in the future
	MongoSession = mongoSession
}

// start it up
func main () {
	startDbSession()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
