package lifts

const (
	TargetLower      = 0
	TargetUpper      = 1
	LiftSquatID      = 2
	LiftOHPID        = 3
	LiftDLID         = 4
	LiftBenchID      = 5
	LiftBarbellRowID = 6
	LiftChinupID     = 8
	LiftPullupIP     = 14
	LiftPushupID     = 59
	LiftHipThrustID  = 81
)

func Bench() Lift {
	return Lift{
		Name:       "Bench Press",
		ExerciseID: LiftBenchID,
		Target:     TargetUpper,
	}
}

func Ohp() Lift {
	return Lift{
		Name:       "Overhead Press",
		ExerciseID: LiftOHPID,
		Target:     TargetUpper,
	}
}

func BarbellRow() Lift {
	return Lift{
		Name:       "Barbell Row",
		ExerciseID: LiftBarbellRowID,
		Target:     TargetUpper,
	}
}

func Squat() Lift {
	return Lift{
		Name:       "Squat",
		ExerciseID: LiftSquatID,
		Target:     TargetLower,
	}
}

func Deadlift() Lift {
	return Lift{
		Name:       "Deadlift",
		ExerciseID: LiftDLID,
		Target:     TargetLower,
	}
}

func Pushup() Lift {
	return Lift{
		Name:       "Push Up",
		ExerciseID: LiftPushupID,
		Target:     TargetUpper,
	}
}

func Chipup() Lift {
	return Lift{
		Name:       "Chinups",
		ExerciseID: LiftChinupID,
		Target:     TargetUpper,
	}
}

func Pullup() Lift {
	return Lift{
		Name:       "Pull Ups",
		ExerciseID: LiftPullupIP,
		Target:     TargetUpper,
	}
}

func HipThrust() Lift {
	return Lift{
		Name:       "Hip Thrust",
		ExerciseID: LiftHipThrustID,
		Target:     TargetLower,
	}
}
