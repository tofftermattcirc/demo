package main

type Exercise struct {
	Id         int       `json:"id,omitempty"`
	Name       string    `json:"name"`
	// Whether the exercise was started by the trainee (true) or not (false)
	IsStarted  bool      `json:"started,omitempty"`
	// The amount of time (in minutes) the trainee has to complete the exercise
	Delta      int       `json:"delta"`
	// The point in time when the exercise will be completed.
	// Set once the exericise has started. 
	Due        int64     `json:"due,omitempty"`
}

type Exercises []Exercise
