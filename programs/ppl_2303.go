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
			goal = sets.Maintain
			dayName = "Strength"
		case 2:
			goal = sets.Maintain
			dayName = "Volume"
		case 3:
			goal = sets.Maintain
			dayName = "RPT"
		case 4:
			goal = sets.Maintain
			dayName = "OneRM"
		case 5:
			goal = sets.Lite
			dayName = "Lite"
		}

		// Instead of thinking how to schedule days, I'm describing the different lift schemes based on the lift
		benchLiftSchemes := []lifts.LiftScheme{
			lifts.FiveByFive,
			lifts.ThreeByEight,
			lifts.RptFourSets,
			lifts.OneRepMaxTest,
		}

		squatLiftSchemes := []lifts.LiftScheme{
			lifts.FiveByFive,
			lifts.ThreeByEight,
			lifts.RptFourSets,
			lifts.OneRepMaxTest,
		}

		deadliftLiftSchemes := []lifts.LiftScheme{
			lifts.ThreeByFive,
			lifts.ThreeByFive,
			lifts.ThreeByFive,
			lifts.OneRepMaxTest,
		}

		ohpLiftSchemes := []lifts.LiftScheme{
			lifts.ThreeByFive, // 3x5 in order to keep workouts short
			lifts.ThreeByEight,
			lifts.RptFourSets,
			lifts.OneRepMaxTest,
		}

		barbellrowLiftSchemes := []lifts.LiftScheme{
			lifts.ThreeByFive, // 3x5 in order to keep workouts short
			lifts.ThreeByEight,
			lifts.RptFourSets,
			lifts.OneRepMaxTest,
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
					getPrimaryLiftByGoal(verticalPull, lifts.ThreeByFive, goal),
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
