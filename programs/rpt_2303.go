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

		// Instead of thinking how to schedule days, I'm describing the different lift schemes based on the lift
		benchLiftSchemes := []lifts.LiftScheme{
			lifts.ThreeByEight,
			lifts.FiveByFive,
			lifts.RptFourSets,
		}

		squatLiftSchemes := []lifts.LiftScheme{
			lifts.FiveByFive,
			lifts.ThreeByEight,
			lifts.RptFourSets,
		}

		deadliftLiftSchemes := []lifts.LiftScheme{
			lifts.ThreeByFive,
			lifts.ThreeByFive,
			lifts.ThreeByFive,
		}

		ohpLiftSchemes := []lifts.LiftScheme{
			lifts.ThreeByFive, // 3x5 in order to keep workouts short
			lifts.ThreeByFive,
			lifts.RptFourSets,
		}

		extraLifts := []lifts.Lift{
			lifts.Deadlift(),
			lifts.Ohp(),
		}

		var extraScheme lifts.LiftScheme

		loopsPerWeek := 3

		for dayNum := 0; dayNum < loopsPerWeek; dayNum++ {
			workoutNum := workoutNum(weeknum, dayNum, loopsPerWeek)

			extraLift := extraLifts[workoutNum%len(extraLifts)]
			extraScheme = ohpLiftSchemes[workoutNum%len(ohpLiftSchemes)]
			extraScheme = deadliftLiftSchemes[workoutNum%len(ohpLiftSchemes)]

			workoutWeek.addWorkoutDay(
				fmt.Sprintf("%s %d", dayName, dayNum+1),
				workoutDay{
					getPrimaryLiftByGoal(lifts.Squat(), squatLiftSchemes[workoutNum%len(squatLiftSchemes)], goal),
					getPrimaryLiftByGoal(lifts.Bench(), benchLiftSchemes[workoutNum%len(benchLiftSchemes)], goal),
					getPrimaryLiftByGoal(extraLift, extraScheme, goal),
				},
			)
		}

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
