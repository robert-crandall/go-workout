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

func setRestTimer(goal sets.Goal) int {
	switch goal {
	case sets.Maintain:
		return 90
	case sets.Increase:
		return 120
	case sets.OneRM:
		return 300
	case sets.Lite:
		return 90
	default:
		return 90
	}
}
