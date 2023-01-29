package programs

import (
	"fmt"
	"main/lifts"
	"main/sets"
)

func init() {
	programs = append(programs, rpt_23())
}

// Workout program designed for squats and bench everyday, deadlift and OHP as accessory lifts
func rpt_23() Program {

	routine := func(weeknum int) ([]string, []workoutDay) {

		var dayNames []string
		var daysList []workoutDay

		var goal sets.Goal

		// Create a new day goal based on the week number
		switch weeknum {
		case 1:
			goal = sets.Maintain
		case 2:
			goal = sets.Increase
		case 3:
			goal = sets.OneRM
		case 4:
			goal = sets.Lite
		}

		weekName := goal.String()

		pullWorkout := alternateLifts([]lifts.Lift{
			lifts.Chipup(),
			lifts.BarbellRow(), // Barbell row in second spot means it will be used on week 3
		}, weeknum, 0, 1)

		extraLifts := []lifts.Lift{
			lifts.Deadlift(),
			lifts.Ohp(),
			pullWorkout,
		}

		// Lift schemes for each day of the week, except 1rm week
		dailyLiftSchemes := []sets.LiftScheme{
			sets.FiveByFive,
			sets.ThreeByEight,
			sets.RptFourSets,
		}

		var liftScheme, secondaryLiftScheme sets.LiftScheme

		switch goal {
		case sets.Maintain, sets.Increase:
			secondaryLiftScheme = sets.ThreeByFive
		case sets.OneRM:
			liftScheme = sets.OneRepMaxTest
			secondaryLiftScheme = sets.OneRepMaxTest
		case sets.Lite:
			secondaryLiftScheme = sets.ThreeByFive
		}

		daysPerWeek := 3
		dayNum := 1

		for dayNum <= daysPerWeek {
			// Allow the 1rm max week to override the daily lift scheme
			if liftScheme != sets.OneRepMaxTest {
				liftScheme = dailyLiftSchemes[dayNum-1]
			}
			extraLift := extraLifts[dayNum-1]

			// No point in having lots of sets on lite week
			if goal == sets.Lite {
				liftScheme = sets.FiveByFive
			}

			var workout workoutDay

			// Standard case - every day workout for each week. Override below.
			workout = workoutDay{
				getPrimaryLift(lifts.Squat(), liftScheme, goal),
				getPrimaryLift(lifts.Bench(), liftScheme, goal),
				getPrimaryLift(extraLift, secondaryLiftScheme, goal),
			}

			dayNames = append(dayNames, fmt.Sprintf("%s %s", weekName, liftScheme.String()))

			// What to do on the 1rm week
			if goal == sets.OneRM {
				switch dayNum {
				case 1:
					workout = workoutDay{
						getPrimaryLift(lifts.Squat(), liftScheme, goal),
						getPrimaryLift(lifts.Bench(), liftScheme, goal),
					}
				case 2:
					workout = workoutDay{
						getPrimaryLift(lifts.Deadlift(), liftScheme, goal),
						getPrimaryLift(lifts.Ohp(), liftScheme, goal),
					}
				case 3:
					goal = sets.Lite // Set the goal to lite for the last day of 1RM week
					liftScheme = sets.FiveByFive
					workout = workoutDay{
						getPrimaryLift(lifts.Squat(), liftScheme, goal),
						getPrimaryLift(lifts.Bench(), liftScheme, goal),
					}
				}
			}

			daysList = append(daysList, workout)
			dayNum = dayNum + 1
		}

		return dayNames, daysList

	}

	program := Program{
		Name:        "RPT 23.01",
		Explanation: "A mix of RPT and 5x5 exercises. Includes OPTIONAL lite lifts for days when things are heavy.",
		Weeks:       4,
		Export:      true,
		Routine:     routine,
	}

	return program
}

// Get a primary workout given the liftScheme and goal
func getPrimaryLift(lift lifts.Lift, liftScheme sets.LiftScheme, goal sets.Goal) workout {
	var thisSets sets.Sets
	var options sets.Options

	options = sets.NewOptions()

	// Set rep count for RPT programs
	if liftScheme.IsRPT() {
		if lift.Target == lifts.TargetLower {
			options = sets.WithRepCount(4)
		} else {
			options = sets.WithRepCount(6)
		}
	}

	thisSets = sets.NewSets(goal, liftScheme, options)

	lastSetIsAmrap := false
	if goal == sets.OneRM {
		lastSetIsAmrap = true
	}

	return workout{
		Lift:           lift,
		IncrementType:  IncrementWeightsOff,
		Sets:           thisSets.GetProgram(),
		LastSetIsAmrap: lastSetIsAmrap,
	}
}
