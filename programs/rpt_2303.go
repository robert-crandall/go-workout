package programs

import (
	"fmt"
	"main/lifts"
	"main/sets"
)

func init() {
	programs = append(programs, rpt_2303())
}

// Workout program designed for squats and bench everyday, deadlift and OHP as accessory lifts
func rpt_2303() Program {

	routine := func(weeknum int) ([]string, []workoutDay) {

		var workoutWeek WorkoutWeek

		var goal sets.Goal
		var dayName string

		// Set weekly goals
		switch weeknum {
		case 1:
			goal = sets.Maintain
			dayName = "Maintain"
		case 2:
			goal = sets.Increase
			dayName = "Increase"
		case 3:
			goal = sets.OneRM
			dayName = "OneRM"
		}

		loopsPerWeek := 3

		for dayNum := 1; dayNum <= loopsPerWeek; dayNum++ {
			workoutNum := workoutNum(weeknum, dayNum, loopsPerWeek)

			firstLift := lifts.Squat()
			secondLift := lifts.Bench()
			extraLift := lifts.Deadlift()

			firstLiftScheme, secondLiftScheme := lifts.FiveByFive, lifts.ThreeByEight
			extraLiftScheme := lifts.ThreeByFive

			// Next week flip them around in order to flip OHP and Deadlift
			if workoutNum%2 == 0 {
				firstLift = lifts.Bench()
				secondLift = lifts.Squat()
				extraLift = lifts.Ohp()
				extraLiftScheme = lifts.ThreeByEight
			}

			// On last day of week, do RPT
			if dayNum == loopsPerWeek {
				firstLiftScheme, secondLiftScheme = lifts.RptFourSets, lifts.RptFourSets
				extraLiftScheme = lifts.RptFourSets
			}

			workoutWeek.addWorkoutDay(
				fmt.Sprintf("%s %d", dayName, dayNum),
				workoutDay{
					getPrimaryLiftByGoal(firstLift, firstLiftScheme, goal),
					getPrimaryLiftByGoal(secondLift, secondLiftScheme, goal),
					getPrimaryLiftByGoal(extraLift, extraLiftScheme, goal),
				},
			)
		}

		// Custom week for One RM testing
		if goal == sets.OneRM {
			workoutWeek.resetWorkoutDays()

			workoutWeek.addWorkoutDay(
				fmt.Sprintf("Squat/Bench One Rep Test"),
				workoutDay{
					getPrimaryLiftByGoal(lifts.Squat(), lifts.OneRepMaxTest, goal),
					getPrimaryLiftByGoal(lifts.Bench(), lifts.OneRepMaxTest, goal),
				},
			)

			workoutWeek.addWorkoutDay(
				fmt.Sprintf("Deadlift/OHP One Rep Test"),
				workoutDay{
					getPrimaryLiftByGoal(lifts.Deadlift(), lifts.OneRepMaxTest, goal),
					getPrimaryLiftByGoal(lifts.Ohp(), lifts.OneRepMaxTest, goal),
				},
			)

			// Throw on the optional lite day. Putting it here instead of another loop so the weekly numbers match.
			recoveryScheme := lifts.ThreeByFive
			recoveryGoal := sets.Lite
			workoutWeek.addWorkoutDay(
				fmt.Sprintf("Recovery Day"),
				workoutDay{
					getPrimaryLiftByGoal(lifts.Squat(), recoveryScheme, recoveryGoal),
					getPrimaryLiftByGoal(lifts.Bench(), recoveryScheme, recoveryGoal),
					getPrimaryLiftByGoal(lifts.Deadlift(), recoveryScheme, recoveryGoal),
					getPrimaryLiftByGoal(lifts.Ohp(), recoveryScheme, recoveryGoal),
				},
			)
		}

		return workoutWeek.getDayNames(), workoutWeek.getWorkouts()
	}

	program := Program{
		Name:        "RPT 23.03",
		Explanation: "A mix of RPT and 5x5 exercises in a 3 day split, focused on Squat and Bench. Includes OPTIONAL lite lifts for days when things are heavy.",
		Weeks:       3,
		Export:      true,
		Routine:     routine,
	}

	return program
}
