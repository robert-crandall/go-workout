package programs

import (
	"main/lifts"
	"main/sets"
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
	Set           sets.Set
	ExerciseType  int
	IncrementType int
}

var (
	programs []Program
)

func GetPrograms() []Program {
	return programs
}

// These values match Personal Trainer app values
const (
	ExerciseTypeWeightBased       = 0  // Increase weight after completed Day
	ExerciseTypeWendlerMainLift   = 1  // Increase weight after completed Program
	ExerciseTypePercentage        = 5  // Increase weight after completed Day
	ExerciseTypeWendlerAssistance = 6  // Do not increase weight
	IncrementTypeYes              = 10 // Follow increase weight rules
	IncrementTypeNo               = 0  // Ignore increase weight rules
)
