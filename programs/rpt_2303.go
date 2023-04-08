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

	routine := func(weeknum int, programSet *ProgramSet) ([]string, []workoutDay) {

		bench := &programSet.Bench
		squat := &programSet.Squat
		deadLift := &programSet.Deadlift
		ohp := &programSet.OHP

		var workoutWeek WorkoutWeek

		var goal sets.Goal
		var dayName string

		// Set weekly goals
		switch weeknum {
		case 1:
			goal = sets.Lite
			dayName = "Lite"
		case 2:
			goal = sets.Maintain
			dayName = "Maintain"
		case 3:
			goal = sets.Increase
			dayName = "Increase"
		case 4:
			goal = sets.OneRM
			dayName = "OneRM"
		}

		bench.AddLiftSchemes([]lifts.LiftScheme{
			lifts.RptFourSets,
		})

		squat.AddLiftSchemes([]lifts.LiftScheme{
			lifts.FiveByFive,
		})

		deadLift.AddLiftSchemes([]lifts.LiftScheme{
			lifts.ThreeByFive,
		})

		ohp.AddLiftSchemes([]lifts.LiftScheme{
			lifts.ThreeByEight,
		})

		loopsPerWeek := 3

		for dayNum := 1; dayNum <= loopsPerWeek; dayNum++ {
			workoutNum := workoutNum(weeknum, dayNum, loopsPerWeek)

			firstLift := squat
			secondLift := bench
			extraLift := deadLift

			// Next week flip them around in order to flip OHP and Deadlift
			if workoutNum%2 == 0 {
				firstLift = bench
				secondLift = squat
				extraLift = ohp
			}

			workoutWeek.addWorkoutDay(
				fmt.Sprintf("%s %d", dayName, dayNum),
				workoutDay{
					getPrimaryLiftByGoal(*firstLift, firstLift.NextScheme(), goal),
					getPrimaryLiftByGoal(*secondLift, secondLift.NextScheme(), goal),
					getPrimaryLiftByGoal(*extraLift, extraLift.NextScheme(), goal),
				},
			)
		}

		// Custom week for One RM testing
		if goal == sets.OneRM {
			workoutWeek.resetWorkoutDays()

			workoutWeek.addWorkoutDay(
				fmt.Sprintf("Squat/Bench One Rep Test"),
				workoutDay{
					getPrimaryLiftByGoal(*squat, lifts.OneRepMaxTest, goal),
					getPrimaryLiftByGoal(*bench, lifts.OneRepMaxTest, goal),
				},
			)

			workoutWeek.addWorkoutDay(
				fmt.Sprintf("Deadlift/OHP One Rep Test"),
				workoutDay{
					getPrimaryLiftByGoal(*deadLift, lifts.OneRepMaxTest, goal),
					getPrimaryLiftByGoal(*ohp, lifts.OneRepMaxTest, goal),
				},
			)

			// Throw on the optional lite day. Putting it here instead of another loop so the weekly numbers match.
			recoveryScheme := lifts.ThreeByFive
			recoveryGoal := sets.Lite
			workoutWeek.addWorkoutDay(
				fmt.Sprintf("Recovery Day"),
				workoutDay{
					getPrimaryLiftByGoal(*squat, recoveryScheme, recoveryGoal),
					getPrimaryLiftByGoal(*bench, recoveryScheme, recoveryGoal),
					getPrimaryLiftByGoal(*deadLift, recoveryScheme, recoveryGoal),
					getPrimaryLiftByGoal(*ohp, recoveryScheme, recoveryGoal),
				},
			)
		}

		return workoutWeek.getDayNames(), workoutWeek.getWorkouts()
	}

	program := Program{
		Name:        "RPT 23.03",
		Explanation: "A mix of RPT and 5x5 exercises in a 3 day split, focused on Squat and Bench. Includes OPTIONAL lite lifts for days when things are heavy.",
		Weeks:       4,
		Export:      true,
		Routine:     routine,
	}

	return program
}
