package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

/** GET /exercises */
func GetExercises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(exercises); err != nil {
		panic(err)
	}
}

/** GET /exercise/{id} */
func GetExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var exerciseId int
	var err error
	if exerciseId, err = strconv.Atoi(vars["id"]); err != nil {
		panic(err)
	}
	exc := RepoGetExercise(exerciseId)
	if exc.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(exc); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Exercise"}' http://localhost:8080/exercises

*/
func AddExercise(w http.ResponseWriter, r *http.Request) {
	var exc Exercise
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, &exc); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	e := RepoAddExercise(exc)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}

func StartExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var exerciseId int
	var err error
	if exerciseId, err = strconv.Atoi(vars["id"]); err != nil {
		panic(err)
	}
	exc := RepoGetExercise(exerciseId)
	if exc.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if !exc.IsStarted {
			exc.IsStarted = true;
			var due = time.Now().Local().Add(time.Duration(exc.Delta) * time.Minute)
			exc.Due = due.UnixNano() / int64(time.Millisecond)
			RepoUpdateExercise(exc)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(exc); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
