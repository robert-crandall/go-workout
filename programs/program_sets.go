package programs

import (
	"main/lifts"
)

// A way to keep track of desired sets and workout numbers for a given lift
type ProgramSet struct {
	Bench    lifts.Lift
	Deadlift lifts.Lift
	Squat    lifts.Lift
	OHP      lifts.Lift
}

func NewProgramSet() ProgramSet {
	return ProgramSet{
		Bench:    lifts.Bench(),
		Deadlift: lifts.Deadlift(),
		Squat:    lifts.Squat(),
		OHP:      lifts.Ohp(),
	}
}
