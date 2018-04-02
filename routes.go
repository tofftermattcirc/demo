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
	// Will return JSON of:
	//    * UserName [String]
	//    * Password [String]  - Unencrypted for POC.
	//    * Validity [boolean]
	// For POC this is using unecrypted info this should be
	// changed later
	Route{
		"AuthenticateUser",
		"POST",
		"/auth/user/{userid}/password/{password}",
		AuthenticateUser,
	},
}
