package programs

import (
	"fmt"
	"main/lifts"
	"main/sets"
)

func init() {
	programs = append(programs, simple_bigFour())
}

// Workout program that is very simple to aim for compliance
func simple_bigFour() Program {

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

		simpleThreeByFive := sets.NewSets(
			sets.WithSetCount(3),
			sets.WithRepCount(5),
			sets.WithWeightPercentage(1),
			sets.WithRestTimer(120),
		).Static()

		simpleThreeByEight := sets.NewSets(
			sets.WithSetCount(3),
			sets.WithRepCount(8),
			sets.WithWeightPercentage(1),
			sets.WithRestTimer(120),
		).Static()

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
			"static": {
				Lift:          lifts.Bench(),
				IncrementType: IncrementWeightsProgramComplete,
				Sets:          simpleThreeByEight,
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
			"static": {
				Lift:          lifts.Ohp(),
				IncrementType: IncrementWeightsProgramComplete,
				Sets:          simpleThreeByEight,
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
			"static": {
				Lift:          lifts.Deadlift(),
				IncrementType: IncrementWeightsProgramComplete,
				Sets:          simpleThreeByFive,
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
				Sets:          simpleThreeByFive.Static(),
				IncrementType: IncrementWeightsPerSession,
			},
			"bodyweight": {
				Lift:          lifts.ReverseLunge(), // It's not really, but I don't want to add a jump squat right now
				Sets:          bodyWeightSet,
				IncrementType: IncrementWeightsProgramComplete,
			},
			"static": {
				Lift:          lifts.Squat(),
				IncrementType: IncrementWeightsProgramComplete,
				Sets:          simpleThreeByEight,
			},
		}

		// Describes WHEN to do exercises
		dayNames = append(dayNames, fmt.Sprintf("Squat %d", weeknum))
		daysList = append(daysList, []workout{
			squat["static"],
			bench["static"],
		})

		dayNames = append(dayNames, fmt.Sprintf("Deadlift %d", weeknum))
		daysList = append(daysList, []workout{
			deadlift["static"],
			ohp["static"],
		})

		return dayNames, daysList

	}

	program := Program{
		Name:        "Simple 22.02",
		Explanation: "Simple workout to encourage compliance",
		Weeks:       2,
		Export:      false,
		Routine:     routine,
	}

	return program
}
