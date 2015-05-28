package main

import (
	//	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CreateNewUser(vars["username"], "password")
}

func Report(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mine := "Do not set mine"
	if vars["setMine"] == "1" {
		mine = "Set mine"
		lat, err := strconv.ParseFloat(vars["lat"], 64)
		lon, err := strconv.ParseFloat(vars["lon"], 64)
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
