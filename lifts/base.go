package lifts

// Lift matches the JSON format of Personal Trainer app
type Lift struct {
	Name       string `json:"name"`
	ExerciseID int    `json:"exercise_id"`
	Target     int    `json:"target"`
}
