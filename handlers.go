package main

import (
//	"encoding/json"
	"fmt"
	"net/http"
	
	"github.com/gorilla/mux"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "Welcome ", vars["username"])
}

func Report(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintln(w, "Latitude: ", vars["lat"], "Longitude: ", vars["lon"])
	
	// todos := Todos{
	// 	Todo{Name: "Write presentation"},
	// 	Todo{Name: "Host meetup"},
	// }
	// if err := json.NewEncoder(w).Encode(todos); err != nil {
	// 	panic(err)
	// }
}
