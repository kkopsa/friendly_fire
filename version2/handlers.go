package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"math"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CreateNewUser(vars["username"], "password")
}

func Report(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mine := "Do not set mine"

	lat, err := strconv.ParseFloat(vars["lat"], 64)
	lon, err := strconv.ParseFloat(vars["lon"], 64)
	
	if (isHit(lat, lon)) {
		fmt.Fprintln(w, "You got hit and now are dead...")
		return
	}

	if vars["setMine"] == "1" {
		mine = "Set mine"
		if err != nil {
			panic("Failed to report mines\n")
		}
		coordinates := []float64{lat, lon}
		SetMine(vars["username"], coordinates)
	}
	fmt.Fprintln(w, "User: ", vars["username"], " \nReporting location: \n\tLatitude: ", vars["lat"], "\n\tLongitude: ", vars["lon"], "\n\t", mine)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := GetUser("not-used-right-now")
	fmt.Fprintln(w, users)
}

func GetMines(w http.ResponseWriter, r *http.Request) {
	mines := GetAllMines()
	fmt.Fprintln(w, mines)
}

func isHit(lat, lon float64) bool {
	//get all mines
	mines := GetAllMines()
	
	earthRadius := 6371000.0 // meters
	
	for _, mine := range mines {
		lat1 := lat
		lon1 := lon
		
		lat2 := mine.Location[0]
		lon2 := mine.Location[1]
		
		dLat := toRad(lat2 - lat1)
		dLon := toRad(lon2 - lon1)
		lat1 = toRad(lat1)
		lat2 = toRad(lat2)
		a := math.Sin(dLat / 2) * math.Sin(dLat / 2) + math.Sin(dLon / 2) * math.Sin(dLon / 2) * math.Cos(lat1) * math.Cos(lat2)
		c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1 - a))
		d := earthRadius * c
		
		if d < 100.0 { // todo change this to a global constant probably...
			return true // Also remember that this is in whatever 
			            // unit of measure that the earth radius is in
		}
	}
	return false
}

func toRad(degrees float64) float64 {
	return (degrees * math.Pi) / 180.0
}
