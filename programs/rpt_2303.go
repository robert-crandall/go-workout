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
		var liftSchemes []sets.LiftScheme

		// Set weekly goals
		switch weeknum {
		case 1:
			goal = sets.Maintain
			liftSchemes = []sets.LiftScheme{
				sets.FiveByFive,   // Monday, Tuesday
				sets.ThreeByEight, // Thursday, Friday
			}
		case 2:
			goal = sets.Increase
			liftSchemes = []sets.LiftScheme{
				sets.RptFourSets,   // Monday, Tuesday
				sets.OneRepMaxTest, // Thursday, Friday
			}
		case 3:
			goal = sets.Lite
			liftSchemes = []sets.LiftScheme{
				sets.FiveByFive, // Monday, Tuesday
				sets.FiveByFive, // Thursday, Friday
			}
		}

		loopsPerWeek := 2

		for dayNum := 1; dayNum <= loopsPerWeek; dayNum++ {
			// Set the lift scheme for the day
			liftScheme := liftSchemes[(dayNum+1)%loopsPerWeek]
			secondaryLiftScheme := sets.ThreeByFive // Default to 3x5 for secondary lifts

			if liftScheme == sets.OneRepMaxTest {
				secondaryLiftScheme = sets.OneRepMaxTest
			}

			// Skip any extra lite sets
			if goal == sets.Lite && dayNum > 1 {
				continue
			}

			verticalPull := alternateLifts([]lifts.Lift{
				lifts.Chipup(),
				lifts.Pullup(),
			}, dayNum, 0, 1)

			// Days 1 & 3, ie Monday and Thursday
			workoutWeek.addWorkoutDay(
				fmt.Sprintf("%s %s %s", "Squat", liftScheme.String(), goal.String()),
				workoutDay{
					getPrimaryLiftByGoal(lifts.Squat(), liftScheme, goal),
					getPrimaryLiftByGoal(lifts.Bench(), liftScheme, goal),
					getPrimaryLiftByGoal(lifts.BarbellRow(), secondaryLiftScheme, goal),
				},
			)

			// Days 2 & 4, ie Tuesday and Friday
			workoutWeek.addWorkoutDay(
				fmt.Sprintf("%s %s %s", "Deadlift", secondaryLiftScheme.String(), goal.String()),
				workoutDay{
					getPrimaryLiftByGoal(lifts.Deadlift(), secondaryLiftScheme, goal),
					getPrimaryLiftByGoal(lifts.Ohp(), liftScheme, goal),
					getPrimaryLiftByGoal(verticalPull, secondaryLiftScheme, goal),
				},
			)
		}

		return workoutWeek.getDayNames(), workoutWeek.getWorkouts()
	}

	program := Program{
		Name:        "RPT 23.03",
		Explanation: "A mix of RPT and 5x5 exercises in a 4 day split. Includes OPTIONAL lite lifts for days when things are heavy.",
		Weeks:       3,
		Export:      true,
		Routine:     routine,
	}

	return program
}
