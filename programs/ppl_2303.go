package programs

import (
	"fmt"
	"main/lifts"
	"main/sets"
)

func init() {
	programs = append(programs, ppl_2303())
}

// Workout program designed for squats and bench everyday, deadlift and OHP as accessory lifts
func ppl_2303() Program {

	routine := func(weeknum int) ([]string, []workoutDay) {

		var workoutWeek WorkoutWeek

		var goal sets.Goal
		var dayName string

		// Set weekly goals
		switch weeknum {
		case 1:
			goal = sets.Increase
			dayName = "Strength"
		case 2:
			goal = sets.Increase
			dayName = "Volume"
		case 3:
			goal = sets.Increase
			dayName = "RPT"
		case 4:
			goal = sets.Increase
			dayName = "OneRM"
		case 5:
			goal = sets.Lite
			dayName = "Lite"
		}

		// Instead of thinking how to schedule days, I'm describing the different lift schemes based on the lift
		benchLiftSchemes := []sets.LiftScheme{
			sets.FiveByFive,
			sets.ThreeByEight,
			sets.RptFourSets,
			sets.OneRepMaxTest,
		}

		squatLiftSchemes := []sets.LiftScheme{
			sets.FiveByFive,
			sets.ThreeByEight,
			sets.RptFourSets,
			sets.OneRepMaxTest,
		}

		deadliftLiftSchemes := []sets.LiftScheme{
			sets.ThreeByFive,
			sets.ThreeByFive,
			sets.ThreeByFive,
			sets.OneRepMaxTest,
		}

		ohpLiftSchemes := []sets.LiftScheme{
			sets.ThreeByFive, // 3x5 in order to keep workouts short
			sets.ThreeByEight,
			sets.RptFourSets,
			sets.OneRepMaxTest,
		}

		barbellrowLiftSchemes := []sets.LiftScheme{
			sets.ThreeByFive, // 3x5 in order to keep workouts short
			sets.ThreeByEight,
			sets.RptFourSets,
			sets.OneRepMaxTest,
		}

		loopsPerWeek := 1

		for dayNum := 1; dayNum <= loopsPerWeek; dayNum++ {
			liftScheme := (weeknum - 1) % len(benchLiftSchemes)

			verticalPull := alternateLifts([]lifts.Lift{
				lifts.Chipup(),
				lifts.Pullup(),
			}, weeknum, 0, 1)

			workoutWeek.addWorkoutDay(
				fmt.Sprintf("Push Day %s", dayName),
				workoutDay{
					getPrimaryLiftByGoal(lifts.Bench(), benchLiftSchemes[liftScheme], goal),
					getPrimaryLiftByGoal(lifts.Ohp(), ohpLiftSchemes[liftScheme], goal),
				},
			)

			workoutWeek.addWorkoutDay(
				fmt.Sprintf("Leg Day %s", dayName),
				workoutDay{
					getPrimaryLiftByGoal(lifts.Squat(), squatLiftSchemes[liftScheme], goal),
					getPrimaryLiftByGoal(lifts.Deadlift(), deadliftLiftSchemes[liftScheme], goal),
				},
			)

			workoutWeek.addWorkoutDay(
				fmt.Sprintf("Pull Day %s", dayName),
				workoutDay{
					getPrimaryLiftByGoal(lifts.BarbellRow(), barbellrowLiftSchemes[liftScheme], goal),
					getPrimaryLiftByGoal(verticalPull, sets.ThreeByFive, goal),
				},
			)
		}

		return workoutWeek.getDayNames(), workoutWeek.getWorkouts()
	}

	program := Program{
		Name:        "PPL 23.03",
		Explanation: "A PPL routine with 2 lifts per day. Designed to be finished in 40 minutes.",
		Weeks:       5,
		Export:      true,
		Routine:     routine,
	}

	return program
}
