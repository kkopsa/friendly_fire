package main
import (
	//"net/http"
	
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/createUser/{username}").Name("CreateUser").HandlerFunc(NewUser)
	router.Methods("GET").Path("/report/{username}").Queries("lat", "{lat}", "lon", "{lon}", "setMine", "{setMine}").Name("Report").HandlerFunc(Report)
	router.Methods("GET").Path("/getUsers").Name("GetUsers").HandlerFunc(GetUsers)
	router.Methods("GET").Path("/getMines").Name("GetMines").HandlerFunc(GetMines)

	return router
}
