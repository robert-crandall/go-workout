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

		// Set weekly goals
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

		// Alternate pulling workouts.
		pullWorkout := alternateLifts([]lifts.Lift{
			lifts.Chipup(),
			lifts.BarbellRow(), // Barbell row in second spot means it will be used weeks 1 and 3
		}, weeknum, 0, 1)

		// The lift to perform after squatting and benching. Easier to think of this as Monday, Wednesday, Friday lifts.
		extraLifts := []lifts.Lift{
			lifts.Deadlift(), // Monday extra lift
			lifts.Ohp(),      // Wednesday extra lift
			pullWorkout,      // Friday extra lift
		}

		// Lift schemes for each day of the week, except 1rm week. Again, easier to think of this as Monday, Wednesday, Friday schemes.
		dailyLiftSchemes := []sets.LiftScheme{
			sets.FiveByFive,   // Monday
			sets.ThreeByEight, // Wednesday
			sets.RptFourSets,  // Friday
		}

		daysPerWeek := 3
		dayNum := 1

		for dayNum <= daysPerWeek {
			// Get the lift schemes for the lifts. Want to keep 3x5 for secondary lifts.
			liftScheme := dailyLiftSchemes[dayNum-1]
			secondaryLiftScheme := sets.ThreeByFive

			// Override lift scheme for 1rm week
			if goal == sets.OneRM {
				liftScheme = sets.OneRepMaxTest
				secondaryLiftScheme = sets.OneRepMaxTest
			}

			// No point in having lots of sets on lite week. Just do 5x5.
			if goal == sets.Lite {
				liftScheme = sets.FiveByFive
			}

			extraLift := extraLifts[dayNum-1]

			// Standard case - every day workout for each week. Override below.
			workout := workoutDay{
				getPrimaryLift(lifts.Squat(), liftScheme, goal),
				getPrimaryLift(lifts.Bench(), liftScheme, goal),
				getPrimaryLift(extraLift, secondaryLiftScheme, goal),
			}

			// What to do on the 1rm week. This is all manually set.
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
						getPrimaryLift(lifts.BarbellRow(), secondaryLiftScheme, goal),
					}
				}
			}

			dayNames = append(dayNames, fmt.Sprintf("%s %s", weekName, liftScheme.String()))
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

	restTime := getRestTimer(goal)

	return workout{
		Lift:            lift,
		IncrementType:   IncrementWeightsOff,
		Sets:            thisSets.GetProgram(),
		LastSetIsAmrap:  lastSetIsAmrap,
		RestTimeSeconds: restTime,
	}
}
