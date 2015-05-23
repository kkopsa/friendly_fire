package main

import (
//	"encoding/json"
	"fmt"
	"net/http"
	
	"github.com/gorilla/mux"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	CreateNewUser(vars["username"], "password")
}

func Report(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintln(w, vars["username"], " reporting location: Latitude: ", vars["lat"], "Longitude: ", vars["lon"])
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := GetUser("tski-user")
	fmt.Fprintln(w, users)
}
