package lifts

// Lift matches the JSON format of Personal Trainer app
type Lift struct {
	Name         string `json:"name"`
	Shortname    string
	ExerciseID   int `json:"exercise_id"`
	Target       int `json:"target"`
	LiftSchemes  []LiftScheme
	WorkoutCount int
}

func (l *Lift) AddLiftSchemes(liftSchemes []LiftScheme) {
	l.LiftSchemes = append(l.LiftSchemes, liftSchemes...)
}

func (l *Lift) AddWorkout() {
	l.WorkoutCount++
}

func (l *Lift) NextScheme() LiftScheme {
	scheme := l.LiftSchemes[l.WorkoutCount%len(l.LiftSchemes)]
	l.WorkoutCount++
	return scheme
}

func (l *Lift) ShortName() string {
	if l.Shortname != "" {
		return l.Shortname
	}
	return l.Name
}
