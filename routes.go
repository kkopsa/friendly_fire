package main
import (
	//"net/http"
	
	"github.com/gorilla/mux"
)

// type Route struct {
// 	Name        string
// 	Method      string
// 	Pattern     string
// 	HandlerFunc http.HandlerFunc
// }

//type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// for _, route := range routes {
	// 	var handler http.Handler
		
	// 	handler = route.HandlerFunc
	// 	handler = Logger(handler, route.Name)
	
	// 	router.
	// 		Methods(route.Method).
	// 		Path(route.Pattern).
	// 		Name(route.Name).
	// 		Handler(handler)
	// }
	router.Methods("GET").Path("/createUser/{username}").Name("CreateUser").HandlerFunc(NewUser)
	router.Methods("GET").Path("/{username/{setMine}}").Queries("lat", "{lat}", "lon", "{lon}").Name("Report").HandlerFunc(Report)

	return router
}

// var routes = Routes{
// 	Route{
// 		"Index",
// 		"GET",
// 		"/createUser/{username}",
// 		NewUser,
// 	},
// 	Route{
// 		"TodoIndex",
// 		"GET",
// 		"/{username}/",
// 		Report,
// 	},
// 	Route{
// 		"TodoShow",
// 		"GET",
// 		"/todos/{todoId}",
// 		TodoShow,
// 	},
//}
