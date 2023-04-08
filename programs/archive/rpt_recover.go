package programs

//
//import (
//	"fmt"
//	"main/lifts"
//	"main/sessions"
//)
//
//func init() {
//	program := recover_rpt()
//	programs = append(programs, program)
//}
//
//// Workout program with one lift (squats) as simple Sets by Reps for recovery
//func recover_rpt() Program {
//
//	routine := func(weeknum int) ([]string, []workoutDay) {
//
//		var dayNames []string
//		var DaysList []workoutDay
//
//		// RPT Bench Day
//		dayNames = append(dayNames, fmt.Sprintf("Bench RPT %d", weeknum))
//		DaysList = append(DaysList, []workout{
//			{
//				Lift:          lifts.Bench(),
//				Session:       sets.RPTIncreaseWeight(weeknum, lifts.Bench().Target),
//				IncrementType: IncrementWeightsProgramComplete,
//			},
//			{
//				Lift:          lifts.Pushup(),
//				Session:       sets.StaticSetsIncreaseReps(3, 8, 2),
//				IncrementType: IncrementWeightsProgramComplete,
//			},
//		})
//
//		// RPT Squat Day, Heavy OHP
//		dayNames = append(dayNames, fmt.Sprintf("Squat RPT %d", weeknum))
//		DaysList = append(DaysList, []workout{
//			{
//				Lift:          lifts.Squat(),
//				Session:       sets.StaticSets(weeknum, 3, 1),
//				IncrementType: IncrementWeightsPerSession,
//			},
//			{
//				Lift:          lifts.Ohp(),
//				Session:       sets.StaticSets(weeknum, 3, 0.85),
//				IncrementType: IncrementWeightsProgramComplete,
//			},
//		})
//
//		// Heavy DL Day
//		dayNames = append(dayNames, fmt.Sprintf("Deadlift Heavy %d", weeknum))
//		DaysList = append(DaysList, []workout{
//			{
//				Lift:          lifts.Deadlift(),
//				Session:       sets.RPTIncreaseWeight(weeknum, lifts.Deadlift().Target),
//				IncrementType: IncrementWeightsProgramComplete,
//			},
//			{
//				Lift:          lifts.Chipup(),
//				Session:       sets.StaticSetsIncreaseReps(3, 5, 1),
//				IncrementType: IncrementWeightsProgramComplete,
//			},
//		})
//
//		// Heavy Bench Day
//		dayNames = append(dayNames, fmt.Sprintf("Bench Heavy %d", weeknum))
//		DaysList = append(DaysList, []workout{
//			{
//				Lift:          lifts.Bench(),
//				Session:       sets.StaticSets(weeknum, 3, 0.85),
//				IncrementType: IncrementWeightsProgramComplete,
//			},
//			{
//				Lift:          lifts.Bench(),
//				Session:       sets.FSL531(weeknum),
//				IncrementType: IncrementWeightsOff,
//			},
//		})
//
//		// Heavy Squat Day, RPT OHP
//		dayNames = append(dayNames, fmt.Sprintf("Squat Heavy %d", weeknum))
//		DaysList = append(DaysList, []workout{
//			{
//				Lift:          lifts.Squat(),
//				Session:       sets.StaticSets(weeknum, 3, 1),
//				IncrementType: IncrementWeightsPerSession,
//			},
//			{
//				Lift:          lifts.Ohp(),
//				Session:       sets.RPTIncreaseWeight(weeknum, lifts.Ohp().Target),
//				IncrementType: IncrementWeightsProgramComplete,
//			},
//		})
//
//		// Heavy Hip Thrust Day
//		dayNames = append(dayNames, fmt.Sprintf("Hip Thrust Heavy %d", weeknum))
//		DaysList = append(DaysList, []workout{
//			{
//				Lift:          lifts.HipThrust(),
//				Session:       sets.RPTIncreaseWeight(weeknum, lifts.HipThrust().Target),
//				IncrementType: IncrementWeightsProgramComplete,
//			},
//			{
//				Lift:          lifts.Pullup(),
//				Session:       sets.StaticSetsIncreaseReps(3, 5, 1),
//				IncrementType: IncrementWeightsProgramComplete,
//			},
//		})
//
//		return dayNames, DaysList
//
//	}
//
//	program := Program{
//		Name:        "RPT 21.06",
//		Explanation: "RPT with Squat Recovery",
//		DaysPerWeek: 6,
//		Weeks:       4,
//		Export:      true,
//		Routine:     routine,
//	}
//
//	return program
//}
