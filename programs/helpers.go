package programs

import "main/lifts"

// Return a lift from the given lifts based on week number
func alertnateWorkouts(workouts []workout, weeknum int) workout {
	return workouts[weeknum%len(workouts)]
}

// Rotate through the lifts based on the day and number
func alternateLifts(lifts []lifts.Lift, weeknum, daynum, daysPerWeek int) lifts.Lift {
	absoluteDay := weeknum*daysPerWeek + daynum
	return lifts[absoluteDay%len(lifts)]
}
