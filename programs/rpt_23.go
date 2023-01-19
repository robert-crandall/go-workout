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

		// Describes WHEN to do exercises
		dayNames = append(dayNames, fmt.Sprintf("OHP %d", weeknum))
		daysList = append(daysList, []workout{
			getSquatDayWeek(1, weeknum),
		})

		dayNames = append(dayNames, fmt.Sprintf("Row %d", weeknum))
		daysList = append(daysList, []workout{
			getSquatDayWeek(2, weeknum),
		})

		dayNames = append(dayNames, fmt.Sprintf("Deadlift %d", weeknum))
		daysList = append(daysList, []workout{
			getSquatDayWeek(3, weeknum),
		})

		return dayNames, daysList

	}

	program := Program{
		Name:        "RPT 23.01",
		Explanation: "A mix of RPT and 5x5 exercises",
		Weeks:       4,
		Export:      true,
		Routine:     routine,
	}

	return program
}

// Get a squat workout given the day and week number
func getSquatDayWeek(dayNum int, weekNum int) workout {
	setCount := 3
	startingReps := 4 + (dayNum - 1)
	rptWeightStartingPercentage := 0.8
	strengthWeekWeightPercentage := 0.85
	resetWeekWeightPercentage := 0.65

	if weekNum == 2 {
		rptWeightStartingPercentage = 0.85
	}

	var squatSets sets.Sets

	switch weekNum {
	case 1, 2:
		squatSets = sets.NewSets(
			sets.WithSetCount(setCount),
			sets.WithRepCount(startingReps),
			sets.WithWeightPercentage(rptWeightStartingPercentage),
		).RPT(2, 0.05)
	case 3: // Simple 5x5 for strength week
		squatSets = sets.NewSets(
			sets.WithSetCount(5),
			sets.WithRepCount(5),
			sets.WithWeightPercentage(strengthWeekWeightPercentage),
			sets.WithRestTimer(120),
		).Static()
	case 4: // Rest Week
		squatSets = sets.NewSets(
			sets.WithSetCount(5),
			sets.WithRepCount(5),
			sets.WithWeightPercentage(resetWeekWeightPercentage),
			sets.WithRestTimer(120),
		).Static()
	}

	return workout{
		Lift:          lifts.Squat(),
		IncrementType: IncrementWeightsOff,
		Sets:          squatSets,
	}
}