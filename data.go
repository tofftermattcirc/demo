package main

import "time"

type Exercise struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
	// Whether the exercise was started by the trainee (true) or not (false)
	IsStarted bool `json:"started,omitempty"`
	// The amount of time (in minutes) the trainee has to complete the exercise
	Delta int `json:"delta"`
	// The point in time when the exercise will be completed.
	// Set once the exericise has started.
	Due int64 `json:"due,omitempty"`
	// The boolean flag that indicates if an exercise has been completed
	//     completed (true) or not (false)
	// * NOTE - BZ * Not sure if this is the correct place to show an exercise
	//               has been completed.  Seems like it should be a combo of
	//               user id and exercise id.
	IsCompleted bool `json:"iscompleted,omitempty"`
	// The point in time when the exercise has been completed.
	// * NOTE - BZ * Note sure if this should be a duration or time of day when finished
	CompletedTime time.Time `json:"completedtime,omitempty"`
}

type Exercises []Exercise

type User struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type Users []User
