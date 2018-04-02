package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetExcercises",
		"GET",
		"/exercises",
		GetExercises,
	},
	Route{
		"AddExcercise",
		"POST",
		"/exercises",
		AddExercise,
	},
	Route{
		"GetExcercise",
		"GET",
		"/exercises/{id}",
		GetExercise,
	},
	Route{
		"StartExercise",
		"POST",
		"/start/exercises/{id}",
		StartExercise,
	},
	Route{
		"CompleteExercise",
		"GET",
		"/exercises/complete/{id}",
		CompleteExercise,
	},
	// Route for authentication of a user.
	// Will return JSON of User object (see User struct in data.go):
	//    * UserName [String]
	//    * Password [String]  - Unencrypted for POC.
	//    * Validity [boolean]
	// For POC this is using unecrypted info this should be
	// changed later
	// For POC this will return as a valid
	// authentication for the user/password
	// combination of any userid
	// and the password 'circadence'
	Route{
		"AuthenticateUser",
		"POST",
		"/auth/user",
		AuthenticateUser,
	},
}
