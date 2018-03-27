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
	if exerciseId, err = strconv.Atoi(vars["Id"]); err != nil {
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

/** GET /exercise/complete/{id} */
func GetCompleteExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var exerciseId int
	var err error
        if exerciseId, err = strconv.Atvoi(vars["Id"]); err != nill {
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
	if err := json.Unmarshal(body, &exc); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	e := RepoAddExercise(exc)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}

func StartExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var exerciseId int
	var err error
	if exerciseId, err = strconv.Atoi(vars["Id"]); err != nil {
		panic(err)
	}
	exc := RepoGetExercise(exerciseId)
	if exc.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if !exc.IsStarted {
			exc.IsStarted = true;
			exc.Due = time.Now().Local().Add(time.Duration(exc.Delta) * time.Minute);
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

func CompleteExercise(w http.ResponseWriter, r *http.Request) {
	// When this function is called it needs to mark the exercise as completed.
        //TODO: not sure what needs to be done to do this.
	//NOTE (BZ): Attempting to base this off StartExercise above.  This may not be 
	// the correct thing to do.
	vars := mux.Vars(r)
	var exerciseId int
	var err error
	if exerciseId, err = strconv.Atoi(vars["Id"]); err != nil {
		panic(err)
	}
	// get the exercise from the repository
	exc := RepoGetExercise(exerciseId)  
	// Error checking to find id
	if exc.Id > 0 {
	    	// the exercise was found.  Now mark it complete.  [bz] need to figure out how...
		
    	}


	// 404 if id not found
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(Http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
