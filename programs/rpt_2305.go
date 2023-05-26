package programs

import (
	"fmt"
	"main/lifts"
	"main/sets"
)

func init() {
	programs = append(programs, rpt_2305())
}

// Workout program designed for squats and bench more often, deadlift and OHP as accessory lifts
func rpt_2305() Program {

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
			// workoutNum := workoutNum(weeknum, dayNum, loopsPerWeek)

			firstLift := squat
			secondLift := bench

			switch dayNum {
			case 2:
				firstLift = deadLift
				secondLift = ohp
			case 4:
				firstLift = barbellRow
				secondLift = pullup
			}

			if goal != sets.OneRM {
				workoutWeek.addWorkoutDay(
					fmt.Sprintf("%s/%s %s %d", firstLift.ShortName(), secondLift.ShortName(), dayName, dayNum),
					workoutDay{
						getPrimaryLiftByGoal(*firstLift, firstLift.NextScheme(), goal),
						getPrimaryLiftByGoal(*secondLift, secondLift.NextScheme(), goal),
					},
				)
			} else {
				// Custom week for One RM testing
				scheme := lifts.OneRepMaxTest
				switch dayNum {
				case 1:
					firstLift = squat
					secondLift = bench
				case 2:
					firstLift = deadLift
					secondLift = ohp
				case 3:
					firstLift = barbellRow
					secondLift = pullup
				case 4: // Rest day
					scheme = lifts.ThreeByFive
					goal = sets.Lite
					firstLift = squat
					secondLift = bench
				}
				workoutWeek.addWorkoutDay(
					fmt.Sprintf("%s/%s One Rep Test", firstLift.ShortName(), bench.ShortName()),
					workoutDay{
						getPrimaryLiftByGoal(*squat, scheme, goal),
						getPrimaryLiftByGoal(*bench, scheme, goal),
					},
				)
			}
		}

		return workoutWeek.getDayNames(), workoutWeek.getWorkouts()
	}

	program := Program{
		Name:        "RPT 23.05",
		Explanation: "A 4 day a week program with a focus on squats and bench press. Deadlift and OHP are accessory lifts.",
		Weeks:       4,
		Export:      true,
		Routine:     routine,
	}

	return program
}
