package main

import (
	"fmt"
)

var currentId int

var exercises Exercises
var users Users

func RepoGetExercise(id int) Exercise {
	for _, ex := range exercises {
		if ex.Id == id {
			return ex
		}
	}
	// return empty Exercise if not found
	return Exercise{}
}

//No Lock on repo : TODO: Add Lock
func RepoAddExercise(ex Exercise) Exercise {
	currentId += 1
	ex.Id = currentId
	ex.IsStarted = false
	exercises = append(exercises, ex)
	return ex
}

func RepoDeleteExercise(id int) error {
	for i, ex := range exercises {
		if ex.Id == id {
			exercises = append(exercises[:i], exercises[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not get Exercise with id of %d to delete", id)
}

func RepoUpdateExercise(exercise Exercise) bool {
	for index, ex := range exercises {
		if ex.Id == exercise.Id {
			exercises[index] = exercise
			return true
		}
	}
	return false
}

/**
Function to authenticate user by User Object
// TODO: for POC acts as a get user command.
*/
//TODO: add lock
func ReopAuthenticateUser(requestedUser User) User {
	/*  Expected code post demo   */
	//	for _, usr := range users {
	//		if usr.Id == requestedUser.Id {
	//			return usr
	//		}
	//	}
	//	// return empty User if not found
	//	return User{}
	/* temp code that looks for user tofftermatt with password circadence */
	// and returns an empty user if not found
	emptyUser := User{}
	tofftermatt := User{
		Name: "tofftermatt",
		Id:   "circadence"}

	if requestedUser.Name == "tofftermatt" {
		if requestedUser.Id == "circadence" {

			// return valid tofftermatt user
			return tofftermatt
		}

		// return invalid
		// currently returns empty user if not found
		return emptyUser
	}

	//  return invalid
	// currently returns empty user if not found
	return emptyUser
}
