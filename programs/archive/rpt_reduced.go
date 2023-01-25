package programs

import (
	"fmt"
	"main/lifts"
	"main/sets"
)

func init() {
	programs = append(programs, rpt_reduced())
}

// Workout program with one lift (squats) as simple Sets by Reps for recovery
func rpt_reduced() Program {

	routine := func(weeknum int) ([]string, []workoutDay) {

		var dayNames []string
		var daysList []workoutDay

		var staticWeight, rptStartingUpperPercentage, rptStartingLowerPercentage float64

		switch weeknum {
		case 1:
			rptStartingUpperPercentage = 0.85
			rptStartingLowerPercentage = 0.9
			staticWeight = 0.75
		case 2:
			rptStartingUpperPercentage = 0.825
			rptStartingLowerPercentage = 0.875
			staticWeight = 0.75
		case 3:
			rptStartingUpperPercentage = 0.85
			rptStartingLowerPercentage = 0.9
			staticWeight = 0.75
		case 4:
			rptStartingUpperPercentage = 0.65
			rptStartingLowerPercentage = 0.65
			staticWeight = 0.65
		}

		secondaryLifts := sets.NewSets(
			sets.WithSetCount(5),
			sets.WithRestTimer(60),
			sets.WithRepCount(5),
			sets.WithWeightPercentage(staticWeight),
		)

		upperRPTLifts := sets.NewSets(
			sets.WithSetCount(3),
			sets.WithRepCount(5),
			sets.WithWeightPercentage(rptStartingUpperPercentage),
		).RPT(2, 0.05)

		squatLifts := sets.NewSets(
			sets.WithSetCount(3),
			sets.WithRepCount(5),
			sets.WithWeightPercentage(1),
		)

		deadliftRPT := sets.NewSets(
			sets.WithSetCount(3),
			sets.WithRepCount(4),
			sets.WithWeightPercentage(rptStartingLowerPercentage),
		).RPT(1, 0.05)

		bodyWeightSet := sets.NewSets(
			sets.WithSetCount(3),
			sets.WithRepCount(8),
			sets.WithRestTimer(20),
			sets.WithWeightPercentage(1),
		).Static()

		// Describes WHAT exercises to do
		bench := map[string]workout{
			"volume": {
				Lift:          lifts.Bench(),
				IncrementType: IncrementWeightsOff,
				Sets:          secondaryLifts.Static(),
			},
			"rpt": {
				Lift:          lifts.Bench(),
				Sets:          upperRPTLifts,
				IncrementType: IncrementWeightsProgramComplete,
			},
			"bodyweight": {
				Lift:          lifts.Pushup(),
				IncrementType: IncrementWeightsOff,
				Sets:          bodyWeightSet,
			},
		}

		deadlift := map[string]workout{
			"volume": {
				Lift:          lifts.Deadlift(),
				IncrementType: IncrementWeightsOff,
				Sets:          secondaryLifts.Static(3),
			},
			"rpt": {
				Lift:          lifts.Deadlift(),
				Sets:          deadliftRPT,
				IncrementType: IncrementWeightsProgramComplete,
			},
		}

		squat := map[string]workout{
			"volume": {
				Lift:          lifts.Squat(),
				IncrementType: IncrementWeightsOff,
				Sets:          secondaryLifts.Static(),
			},
			"rpt": {
				Lift:          lifts.Squat(),
				Sets:          squatLifts.Static(),
				IncrementType: IncrementWeightsProgramComplete,
			},
			"bodyweight": {
				Lift:          lifts.ReverseLunge(), // It's not really, but I don't want to add a jump squat right now
				Sets:          bodyWeightSet,
				IncrementType: IncrementWeightsProgramComplete,
			},
		}

		pullup := map[string]workout{
			"bodyweight": {
				Lift:          lifts.Pullup(),
				IncrementType: IncrementWeightsOff,
				Sets:          bodyWeightSet,
			},
		}

		chinup := map[string]workout{
			"bodyweight": {
				Lift:          lifts.Chipup(),
				IncrementType: IncrementWeightsOff,
				Sets:          bodyWeightSet,
			},
		}

		// Describes WHEN to do exercises
		dayNames = append(dayNames, fmt.Sprintf("Bench and Squat %d", weeknum))
		daysList = append(daysList, []workout{
			squat["volume"],
			bench["rpt"],
			bench["bodyweight"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Deadlift Volume %d", weeknum))
		daysList = append(daysList, []workout{
			deadlift["volume"],
			pullup["bodyweight"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Bench and Squat %d", weeknum))
		daysList = append(daysList, []workout{
			bench["volume"],
			squat["rpt"],
			squat["bodyweight"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Deadlift Heavy %d", weeknum))
		daysList = append(daysList, []workout{
			deadlift["rpt"],
			chinup["bodyweight"],
		})

		return dayNames, daysList

	}

	program := Program{
		Name:        "Reduced RPT 21.08",
		Explanation: "Reduced RPT exercise aiming 4x week",
		Weeks:       1,
		Export:      false,
		Routine:     routine,
	}

	return program
}
