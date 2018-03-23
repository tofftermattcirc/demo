package main

import (
	"fmt"
)

var currentId int

var excercises Excercises

func RepoGetExcercise(id int) Excercise {
	for _, ex := range excercises {
		if ex.Id == id {
			return ex
		}
	}
	// return empty Excercise if not found
	return Excercise{}
}

//No Lock on repo : TODO: Add Lock
func RepoAddExcercise(ex Excercise) Excercise {
	currentId += 1
	ex.Id = currentId
	excercises = append(excercises, ex)
	return ex
}

func RepoDeleteExcercise(id int) error {
	for i, ex := range excercises {
		if ex.Id == id {
			excercises = append(excercises[:i], excercises[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not get Exercise with id of %d to delete", id)
}
