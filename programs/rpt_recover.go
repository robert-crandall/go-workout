package programs

import (
	"fmt"
	"main/lifts"
	"main/sets"
)

func init() {
	program := recover_rpt()
	programs = append(programs, program)
}

// Workout program with one lift (squats) as simple Sets by Reps for recovery
func recover_rpt() Program {

	routine := func(weeknum int) ([]string, []workoutDay) {

		var dayNames []string
		var DaysList []workoutDay

		// RPT Bench Day
		dayNames = append(dayNames, fmt.Sprintf("Bench RPT %d", weeknum))
		DaysList = append(DaysList, []workout{{
			Lift:          lifts.Bench(),
			Set:           sets.RPTIncreaseWeight(weeknum, true),
			ExerciseType:  ExerciseTypeWendlerMainLift,
			IncrementType: IncrementTypeYes,
		},
			{
				Lift:          lifts.Pushup(),
				Set:           sets.StaticSetsIncreaseReps(3, 6, 2),
				ExerciseType:  ExerciseTypeWendlerMainLift,
				IncrementType: IncrementTypeYes,
			},
		})

		// Squat Day
		dayNames = append(dayNames, fmt.Sprintf("TEST Squat RPT %d", weeknum))
		DaysList = append(DaysList, []workout{{
			Lift: lifts.Squat(),
			Set:  sets.RPTIncreaseWeight(weeknum, true),
		},
			{
				Lift: lifts.Ohp(),
				Set:  sets.StaticSetsIncreaseReps(3, 6, 2),
			},
		})

		return dayNames, DaysList

	}

	program := Program{
		Name:        "RPT 21.07",
		Explanation: "RPT with Squat Recovery",
		DaysPerWeek: 2,
		Weeks:       4,
		Export:      true,
		Routine:     routine,
	}

	return program
}
