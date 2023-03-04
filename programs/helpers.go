package programs

import (
	"main/lifts"
	"main/sets"
)

// Return a lift from the given lifts based on week number
func alertnateWorkouts(workouts []workout, weeknum int) workout {
	return workouts[weeknum%len(workouts)]
}

// Rotate through the lifts based on the day and number
// To simply alternate, set daysPerWeek to 1 and daynum to 0
func alternateLifts(lifts []lifts.Lift, weeknum, daynum, daysPerWeek int) lifts.Lift {
	absoluteDay := weeknum*daysPerWeek + daynum
	return lifts[absoluteDay%len(lifts)]
}

func getRestTimer(goal sets.Goal) int {
	switch goal {
	case sets.Maintain:
		return 90
	case sets.Increase:
		return 120
	case sets.OneRM:
		return 180
	case sets.Lite:
		return 90
	default:
		return 90
	}
}

// Get a primary workout given the liftScheme and goal
func getPrimaryLiftByGoal(lift lifts.Lift, liftScheme sets.LiftScheme, goal sets.Goal) workout {
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
