package programs

import (
	"fmt"
	"main/lifts"
	"main/sets"
)

func init() {
	program := texas_monthly()
	programs = append(programs, program)
}

// Workout program with one lift (squats) as simple Sets by Reps for recovery
func texas_monthly() Program {

	routine := func(weeknum int) ([]string, []workoutDay) {

		var dayNames []string
		var DaysList []workoutDay

		bench := map[string]workout{
			"recovery": {
				Lift:          lifts.Bench(),
				Set:           sets.StaticSetsIncreaseWeekly(weeknum, 3, .9),
				IncrementType: IncrementWeightsOff,
			},
			"volume": {
				Lift:          lifts.Bench(),
				Set:           sets.StaticSetsIncreaseWeekly(weeknum, 5, .9),
				IncrementType: IncrementWeightsOff,
			},
			"intensity": {
				Lift:          lifts.Bench(),
				Set:           sets.StaticSetsIncreaseWeekly(weeknum, 3, 1),
				IncrementType: IncrementWeightsProgramComplete,
			},
		}

		ohp := map[string]workout{
			"recovery": {
				Lift:          lifts.Ohp(),
				Set:           sets.StaticSetsIncreaseWeekly(weeknum, 3, .9),
				IncrementType: IncrementWeightsOff,
			},
			"volume": {
				Lift:          lifts.Ohp(),
				Set:           sets.StaticSetsIncreaseWeekly(weeknum, 5, .9),
				IncrementType: IncrementWeightsOff,
			},
			"intensity": {
				Lift:          lifts.Ohp(),
				Set:           sets.StaticSetsIncreaseWeekly(weeknum, 3, 1),
				IncrementType: IncrementWeightsProgramComplete,
			},
		}

		// Still recovering squat. Planned sessions are when recovery is done.
		squat := map[string]workout{
			"recovery": {
				Lift:          lifts.Squat(),
				Set:           sets.StaticSets(1, 3, 1),
				IncrementType: IncrementWeightsPerSession,
			},
			"planned_recovery": {
				Lift:          lifts.Squat(),
				Set:           sets.StaticSets(weeknum, 2, .8),
				IncrementType: IncrementWeightsOff,
			},
			"volume": {
				Lift:          lifts.Squat(),
				Set:           sets.StaticSets(1, 3, 1),
				IncrementType: IncrementWeightsPerSession,
			},
			"planned_volume": {
				Lift:          lifts.Squat(),
				Set:           sets.StaticSets(weeknum, 5, .9),
				IncrementType: IncrementWeightsOff,
			},
			"intensity": {
				Lift:          lifts.Squat(),
				Set:           sets.StaticSets(1, 3, 1),
				IncrementType: IncrementWeightsPerSession,
			},
			"planned_intensity": {
				Lift:          lifts.Squat(),
				Set:           sets.StaticSets(weeknum, 3, 1),
				IncrementType: IncrementWeightsProgramComplete,
			},
		}

		deadlift := map[string]workout{
			"recovery": {
				Lift:          lifts.Deadlift(),
				Set:           sets.StaticSets(weeknum, 2, .8),
				IncrementType: IncrementWeightsPerSession,
			},
			"volume": {
				Lift:          lifts.Deadlift(),
				Set:           sets.StaticSets(weeknum, 3, .9),
				IncrementType: IncrementWeightsPerSession,
			},
			"intensity": {
				Lift:          lifts.Deadlift(),
				Set:           sets.StaticSets(weeknum, 1, 1),
				IncrementType: IncrementWeightsProgramComplete,
			},
		}

		barbellRow := map[string]workout{
			"recovery": {
				Lift:          lifts.BarbellRow(),
				Set:           sets.StaticSetsIncreaseWeekly(weeknum, 3, .9),
				IncrementType: IncrementWeightsOff,
			},
			"volume": {
				Lift:          lifts.BarbellRow(),
				Set:           sets.StaticSetsIncreaseWeekly(weeknum, 5, .9),
				IncrementType: IncrementWeightsOff,
			},
			"intensity": {
				Lift:          lifts.BarbellRow(),
				Set:           sets.StaticSetsIncreaseWeekly(weeknum, 3, 1),
				IncrementType: IncrementWeightsProgramComplete,
			},
		}

		dayNames = append(dayNames, fmt.Sprintf("Bench Volume %d", weeknum))
		DaysList = append(DaysList, []workout{
			bench["volume"],
			ohp["intensity"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Squat Heavy %d", weeknum))
		DaysList = append(DaysList, []workout{
			squat["intensity"],
			barbellRow["recovery"],
		})

		verticalPull := lifts.Pullup()
		if weeknum%2 == 0 {
			verticalPull = lifts.Chipup()
		}

		dayNames = append(dayNames, fmt.Sprintf("OHP Volume %d", weeknum))
		DaysList = append(DaysList, []workout{
			ohp["volume"],
			{
				Lift:          verticalPull,
				Set:           sets.StaticSetsIncreaseReps(3, 5, 1),
				IncrementType: IncrementWeightsProgramComplete,
			},
		})

		dayNames = append(dayNames, fmt.Sprintf("Squat Volume %d", weeknum))
		DaysList = append(DaysList, []workout{
			squat["volume"],
			barbellRow["intensity"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Bench Heavy %d", weeknum))
		DaysList = append(DaysList, []workout{
			bench["intensity"],
			bench["recovery"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Squat Recovery %d", weeknum))
		DaysList = append(DaysList, []workout{
			squat["recovery"],
			deadlift["intensity"],
		})

		return dayNames, DaysList

	}

	program := Program{
		Name:        "Texas Monthly 21.07",
		Explanation: "Texas method with monthly progression. Still recovering squat.",
		DaysPerWeek: 6,
		Weeks:       4,
		Export:      true,
		Routine:     routine,
	}

	return program
}
