package programs

import (
	"main/lifts"
	"main/sessions"
)

type Program struct {
	Name        string
	Explanation string
	DaysPerWeek int                                // Workouts per week
	Weeks       int                                // Weeks until a complete cycle
	Export      bool                               // Whether to export this program or not
	Routine     func(int) ([]string, []workoutDay) // Lists of day names and workouts
}

type workoutDay []workout

type workout struct {
	Lift          lifts.Lift
	Session       sessions.Session
	IncrementType int
}

var (
	programs []Program
)

func GetPrograms() []Program {
	return programs
}

// How often to add weight to the lift
const (
	IncrementWeightsOff             = 0
	IncrementWeightsProgramComplete = 1
	IncrementWeightsPerSession      = 2
)
