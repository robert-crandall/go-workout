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
		pullup := &programSet.Pullup
		barbellRow := &programSet.BarbellRow

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
			lifts.RptThreeSets,
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

		pullup.AddLiftSchemes([]lifts.LiftScheme{
			lifts.ThreeByFive,
		})

		barbellRow.AddLiftSchemes([]lifts.LiftScheme{
			lifts.ThreeByFive,
		})

		loopsPerWeek := 4

		for dayNum := 1; dayNum <= loopsPerWeek; dayNum++ {
			workoutNum := workoutNum(weeknum, dayNum, loopsPerWeek)

			firstLift := squat
			secondLift := bench
			extraLift := barbellRow

			// Next week flip them around in order to flip OHP and Deadlift
			if workoutNum%2 == 0 {
				firstLift = deadLift
				secondLift = ohp
				extraLift = pullup
			}

			if goal != sets.OneRM {
				workoutWeek.addWorkoutDay(
					fmt.Sprintf("%s/%s %s %d", firstLift.Name, secondLift.Name, dayName, dayNum),
					workoutDay{
						getPrimaryLiftByGoal(*firstLift, firstLift.NextScheme(), goal),
						getPrimaryLiftByGoal(*secondLift, secondLift.NextScheme(), goal),
						getPrimaryLiftByGoal(*extraLift, extraLift.NextScheme(), goal),
					},
				)
			} else {
				// Custom week for One RM testing
				// Real testing
				if dayNum <= 2 {
					workoutWeek.addWorkoutDay(
						fmt.Sprintf("%s/%s One Rep Test", firstLift.Name, secondLift.Name),
						workoutDay{
							getPrimaryLiftByGoal(*firstLift, lifts.OneRepMaxTest, goal),
							getPrimaryLiftByGoal(*secondLift, lifts.OneRepMaxTest, goal),
							getPrimaryLiftByGoal(*extraLift, lifts.OneRepMaxTest, goal),
						},
					)
				} else {
					// Rest days
					recoveryScheme := lifts.ThreeByFive
					recoveryGoal := sets.Lite
					workoutWeek.addWorkoutDay(
						fmt.Sprintf("%s/%s Recovery", firstLift.Name, secondLift.Name),
						workoutDay{
							getPrimaryLiftByGoal(*firstLift, recoveryScheme, recoveryGoal),
							getPrimaryLiftByGoal(*secondLift, recoveryScheme, recoveryGoal),
							getPrimaryLiftByGoal(*extraLift, recoveryScheme, recoveryGoal),
						},
					)
				}
			}
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
