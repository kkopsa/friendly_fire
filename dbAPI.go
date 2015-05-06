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

	Config struct {
		URL string
		DbName string
		Tables map[string]string
	}

	// Coordinates struct {
	// 	Latitude float64
	// 	Longitude float64
	// }
	
	// User model
	User struct {
		Username           string        `bson:"username"`
		ID                 bson.ObjectId `bson:"_id,omitempty"`
		SaltedPass         string        `bson:"salted_pass"`
		WarInvitations     []string      `bson:"war_invites"`
		ContractOfWar      string        `bson:"contract_of_war"`
		PrevLocation       []float64     `bson:"coordinates"`
		FriendInvites      []string      `bson:"friend_invites"`
		InvitationsSent    []string      `bson:"friend_invites_sent"`
		WarInvitationsSent []string      `bson:"war_invites_sent"`
	}

	// Mine 
	Mine struct {
		ID          bson.ObjectId `bson:"_id,omitempty"`
		Location    []float64     `bson:"coordinates"`
		BlastRadius int32         `bson:"blast_radius"`
		OwnerId     string        `bson:"owner_id"`
	}

	// War
	ContractOfWar struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		RedTeam   []string      `bson:"red_team"`
		BlueTeam  []string      `bson:"blue_team"`
		RedMines  []Mine        `bson:"red_mines"`
		BlueMines []Mine        `bson:"blue_mines"`
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

// func connect() *Collection {
// 	session, err := mgo.Dial(db.URL)
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	session.SetMode(mgo.Monotonic, true)
// 	return session.DB(db.DbName).C(db.Tables[""])
// }

func CreateNewUser(username, password string, waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {
	// Decrement the wait group count so the program knows this
	// has been completed once the goroutine exits.
	defer waitGroup.Done()
 
	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := mongoSession.Copy()
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
	
	// log.Printf("RunQuery : %d : Executing\n", query)
 
	// // Retrieve the list of stations.
	// var buoyStations []BuoyStation
	// err := collection.Find(nil).All(&buoyStations)
	// if err != nil {
	// 	log.Printf("RunQuery : ERROR : %s\n", err)
	// 	return
	// }
 
	// log.Printf("RunQuery : %d : Count[%d]\n", query, len(buoyStations))
	
}

func CreateWar(username string, waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {

	// Decrement the wait group count so the program knows this
	// has been completed once the goroutine exits.
	defer waitGroup.Done()
 
	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
 
	// Get a collection to execute the query against.
	collection := sessionCopy.DB("FriendlyFire").C("wars")
 
	war := ContractOfWar{}
	war.RedTeam = []string{username}
	war.BlueTeam = []string{"SomeGuy"}

	err := collection.Insert(war)
	if err != nil {
		log.Printf("CreateWar : ERROR : %s\n", err)
		return
	}

	log.Printf("CreateWar : created war : %s\n", war.ID)
}


func SetMine(username string, coordinates float64, 
	          waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {

	// Decrement the wait group count so the program knows this
	// has been completed once the goroutine exits.
	defer waitGroup.Done()
 
	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
 
	// Get a collection to execute the query against.
	collection := sessionCopy.DB("FriendlyFire").C("wars")
 
	mine := Mine{}
	mine.Location = []float64{12.212312, 23.12312}
	
	err := collection.Insert(mine)
	if err != nil {
		//log.Fatal(err)
		log.Printf("SetMine : ERROR : %s\n", err)
		return
	}

	log.Printf("SetMine : created mine : %s\n", username)
}


func GetUser(username string, waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {

	// Decrement the wait group count so the program knows this
	// has been completed once the goroutine exits.
	defer waitGroup.Done()
 
	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
 
	// Get a collection to execute the query against.
	collection := sessionCopy.DB("FriendlyFire").C("users")

	var users []User
	err := collection.Find(nil).All(&users)
	if err != nil {
		//log.Fatal(err)
		log.Printf("SetMine : ERROR : %s\n", err)
		return
	}

	log.Printf("SetMine : created mine : %s\n", username)
}


