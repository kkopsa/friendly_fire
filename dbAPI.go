package main

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"encoding/json"
	"sync"
)

const (
	ConfigFile string = "dbConfig.json"
)

type (
	
	MgoDB struct {
		DInfo mgo.DialInfo
		DSession *mgo.Session
		DWaitGroup sync.WaitGroup
	}

	Config struct {
		URL string
		DbName string
		Tables map[string]string
	}
	
	User struct {
		Username           string        `bson:"username"`
		ID                 bson.ObjectId `bson:"_id,omitempty"`
		SaltedPass         string        `bson:"salted_pass"`
		PrevLocation       []float64     `bson:"coordinates"`
	}

	Mine struct {
		ID          bson.ObjectId `bson:"_id,omitempty"`
		Location    []float64     `bson:"coordinates"`
		OwnerId     string        `bson:"owner_id"`
		Status      bool          `bson:"status"`
	}
)

// Returns database struct with necessary credentials to connect to database
func getConfig() Config {
	file, _ := os.Open(ConfigFile)
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return config
}

func CreateNewUser(username, password string) {
 
	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := MongoSession.Copy()
	defer sessionCopy.Close()
 
	// Get a collection to execute the query against.
	collection := sessionCopy.DB("FriendlyFire").C("users")
 
	user := User{}
	user.Username = username
	user.SaltedPass = password
	err := collection.Insert(user)
	if err != nil {
		//log.Fatal(err)
		log.Printf("CreateUser : ERROR : %s\n", err)
		return
	}

	log.Printf("CreateUser : created user : %s\n", username)	
}

func SetMine(username string, coordinates []float64) {
 
	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := MongoSession.Copy()
	defer sessionCopy.Close()
 
	// Get a collection to execute the query against.
	collection := sessionCopy.DB("FriendlyFire").C("mines")
 
	mine := Mine{}
	mine.Location = []float64{coordinates[0], coordinates[1]}
	mine.OwnerId = username
	mine.Status = true
	
	err := collection.Insert(mine)
	if err != nil {
		//log.Fatal(err)
		log.Printf("SetMine : ERROR : %s\n", err)
		return
	}

	log.Printf("SetMine : created mine : %s\n", username)
}


func GetUser(username string) []User {
	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := MongoSession.Copy()
	defer sessionCopy.Close()
 
	// Get a collection to execute the query against.
	collection := sessionCopy.DB("FriendlyFire").C("users")

	var users []User
	err := collection.Find(nil).All(&users)
	if err != nil {
		//log.Fatal(err)
		log.Printf("GetUsers : ERROR : %s\n", err)
		panic("Could not get users")
	}
	log.Printf(username)
	log.Printf("GetUsers : retrieved users : %s\n", users)
	return users
}

func GetAllMines() []Mine {
	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := MongoSession.Copy()
	defer sessionCopy.Close()
 
	// Get a collection to execute the query against.
	collection := sessionCopy.DB("FriendlyFire").C("mines")

	var mines []Mine
	err := collection.Find(nil).All(&mines)
	if err != nil {
		//log.Fatal(err)
		log.Printf("GetMines : ERROR : %s\n", err)
		panic("Could not get mines")
	}
	log.Printf("Mine: %s", mines[len(mines) - 1])
	log.Printf("GetMines : retrieved mines : %s\n", mines)
	return mines
}


