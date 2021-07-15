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
		var daysList []workoutDay

		var staticWeight, rptStartingUpperPercentage, rptStartingLowerPercentage float64
		var verticalPullLift lifts.Lift

		switch weeknum {
		case 1:
			rptStartingUpperPercentage = 0.8
			rptStartingLowerPercentage = 0.85
			staticWeight = 0.75
			verticalPullLift = lifts.Pullup()
		case 2:
			rptStartingUpperPercentage = 0.825
			rptStartingLowerPercentage = 0.875
			staticWeight = 0.75
			verticalPullLift = lifts.Chipup()
		case 3:
			rptStartingUpperPercentage = 0.85
			rptStartingLowerPercentage = 0.9
			staticWeight = 0.75
			verticalPullLift = lifts.Pullup()
		case 4:
			rptStartingUpperPercentage = 0.65
			rptStartingLowerPercentage = 0.65
			staticWeight = 0.65
			verticalPullLift = lifts.Chipup()
		}

		_ = rptStartingLowerPercentage

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
			sets.WithSetCount(1),
			sets.WithRepCount(8),
			sets.WithWeightPercentage(1),
		).Static()

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
				Lift:          lifts.OverheadTriceps(), // Not really. Doing handstands or something.
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
			"bodyweight": {
				Lift:          verticalPullLift,
				IncrementType: IncrementWeightsProgramComplete,
				Sets:          bodyWeightSet,
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
				Lift:          lifts.OneLeggedSquat(), // It's not really, but I don't want to add a jump squat right now
				Sets:          bodyWeightSet,
				IncrementType: IncrementWeightsProgramComplete,
			},
		}

		dayNames = append(dayNames, fmt.Sprintf("Bench %d", weeknum))
		daysList = append(daysList, []workout{
			squat["volume"],
			bench["rpt"],
			bench["bodyweight"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Deadlift %d", weeknum))
		daysList = append(daysList, []workout{
			ohp["volume"],
			deadlift["rpt"],
			deadlift["bodyweight"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Squat %d", weeknum))
		daysList = append(daysList, []workout{
			bench["volume"],
			squat["rpt"],
			squat["bodyweight"],
		})

		dayNames = append(dayNames, fmt.Sprintf("OHP %d", weeknum))
		daysList = append(daysList, []workout{
			deadlift["volume"],
			ohp["rpt"],
			ohp["bodyweight"],
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
