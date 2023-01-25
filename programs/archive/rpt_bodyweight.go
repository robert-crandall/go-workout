package programs

import (
	"fmt"
	"main/lifts"
	"main/sets"
)

func init() {
	programs = append(programs, rpt_bodyweight())
}

// Workout program with one lift (squats) as simple Sets by Reps for recovery
func rpt_bodyweight() Program {

	routine := func(weeknum int) ([]string, []workoutDay) {

		var dayNames []string
		var daysList []workoutDay

		var staticWeight, rptStartingUpperPercentage, rptStartingLowerPercentage float64

		switch weeknum {
		case 1:
			rptStartingUpperPercentage = 0.8
			rptStartingLowerPercentage = 0.85
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
				IncrementType: IncrementWeightsProgramComplete,
				Sets:          bodyWeightSet,
			},
		}

		ohp := map[string]workout{
			"volume": {
				Lift:          lifts.Ohp(),
				IncrementType: IncrementWeightsOff,
				Sets:          secondaryLifts.Static(),
			},
			"rpt": {
				Lift:          lifts.Ohp(),
				Sets:          upperRPTLifts,
				IncrementType: IncrementWeightsProgramComplete,
			},
			"bodyweight": {
				Lift:          lifts.HandstandPushup(), // Not really. Doing handstands or something.
				IncrementType: IncrementWeightsProgramComplete,
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
				IncrementType: IncrementWeightsPerSession,
			},
			"bodyweight": {
				Lift:          lifts.ReverseLunge(), // It's not really, but I don't want to add a jump squat right now
				Sets:          bodyWeightSet,
				IncrementType: IncrementWeightsProgramComplete,
			},
		}

		barbellRow := map[string]workout{
			"volume": {
				Lift:          lifts.BarbellRow(),
				IncrementType: IncrementWeightsOff,
				Sets:          secondaryLifts.Static(),
			},
			"rpt": {
				Lift:          lifts.BarbellRow(),
				Sets:          upperRPTLifts,
				IncrementType: IncrementWeightsProgramComplete,
			},
			// Bodyweight rows would be something like suspended ring rows.
			//"bodyweight": {
			//	Lift:          verticalPullLift,
			//	IncrementType: IncrementWeightsProgramComplete,
			//	Sets:          bodyWeightSet,
			//},
		}

		pullup := map[string]workout{
			"bodyweight": {
				Lift:          lifts.Pullup(),
				IncrementType: IncrementWeightsProgramComplete,
				Sets:          bodyWeightSet,
			},
		}

		chinup := map[string]workout{
			"bodyweight": {
				Lift:          lifts.Chipup(),
				IncrementType: IncrementWeightsProgramComplete,
				Sets:          bodyWeightSet,
			},
		}

		// Describes WHEN to do exercises
		dayNames = append(dayNames, fmt.Sprintf("Push Bench %d", weeknum))
		daysList = append(daysList, []workout{
			ohp["volume"],
			bench["rpt"],
			bench["bodyweight"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Pull Heavy %d", weeknum))
		daysList = append(daysList, []workout{
			barbellRow["rpt"],
			pullup["bodyweight"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Legs Squat %d", weeknum))
		daysList = append(daysList, []workout{
			deadlift["volume"],
			squat["rpt"],
			squat["bodyweight"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Push OHP %d", weeknum))
		daysList = append(daysList, []workout{
			bench["volume"],
			ohp["rpt"],
			ohp["bodyweight"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Pull Volume %d", weeknum))
		daysList = append(daysList, []workout{
			barbellRow["volume"],
			chinup["bodyweight"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Legs Deadlift %d", weeknum))
		daysList = append(daysList, []workout{
			squat["volume"],
			deadlift["rpt"],
		})

		return dayNames, daysList

	}

	program := Program{
		Name:        "RPT 21.07",
		Explanation: "RPT exercises with bodyweight at end. Still recovering squat.",
		Weeks:       4,
		Export:      true,
		Routine:     routine,
	}

	return program
}
