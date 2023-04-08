package exporters

import (
	"encoding/json"
	"log"
	"main/programs"
)

const (
	ExerciseTypeWeightBased       = 0
	ExerciseTypeWendlerMainLift   = 1 // Monthly progression with Warmup
	ExerciseTypePercentage        = 5
	ExerciseTypeWendlerAssistance = 6 // No progression, no warmup
	IncrementTypeYes              = 10
	IncrementTypeNo               = 0
)

type PTAWorkout struct {
	ThemeColor          int          `json:"theme_color"`
	Duration            int          `json:"duration"`
	Days                int          `json:"days"`
	DaysList            [][]DaysList `json:"days_list"`
	Noofdays            int          `json:"noofdays"`
	Explanation         string       `json:"explanation"`
	ShortName           string       `json:"short_name"`
	Routinetype         string       `json:"routinetype"`
	DayNames            []string     `json:"dayNames"`
	Percentage          int          `json:"percentage"`
	Day                 int          `json:"day"`
	WhenToUpdateWeights int          `json:"when_to_update_weights"`
	ProgramDays         int          `json:"program_days"`
	Name                string       `json:"name"`
	Category            int          `json:"category"`
	Date                int          `json:"date"`
	Realdays            int          `json:"realdays"`
	Expanded            []bool       `json:"expanded"`
	Noofexercises       int          `json:"noofexercises"`
}

type DaysList struct {
	Failuresallowed      int       `json:"failuresallowed"`
	Currentweightlb      float64   `json:"currentweightlb"`
	WarmupType           int       `json:"warmup_type"`
	Resttime2            int       `json:"resttime2"`
	DeloadPercentage     int       `json:"deload_percentage"`
	Currentweightkg      float64   `json:"currentweightkg"`
	AllsetsareAMRAP      bool      `json:"allsetsareAMRAP"`
	SameWeight           bool      `json:"same_weight"`
	BarbellKg            int       `json:"barbell_kg"`
	BarbellLb            int       `json:"barbell_lb"`
	Resttime3            int       `json:"resttime3"`
	Incrementkg          int       `json:"incrementkg"`
	Type                 int       `json:"type"`
	Increment            string    `json:"increment"`
	Failures             int       `json:"failures"`
	SameWeightWarmUp     bool      `json:"same_weight_warm_up"`
	Doubleincrement      bool      `json:"doubleincrement"`
	DependentIncrementID int       `json:"dependentIncrementID"`
	DependentExerciseID  int       `json:"dependent_exercise_id"`
	Failure              int       `json:"failure"`
	ResttimeWarmup       int       `json:"resttime_warmup"`
	Resttime1            int       `json:"resttime1"`
	SupersetNumber       int       `json:"supersetNumber"`
	WeightskgList        []float64 `json:"weightskg_list"`
	RepsList             []int     `json:"reps_list"`
	LastsetsisAMRAP      bool      `json:"lastsetsisAMRAP"`
	IncrementType        int       `json:"increment_type"`
	WeightslbList        []int     `json:"weightslb_list"`
	Incrementlb          int       `json:"incrementlb"`
	PercentageList       []float64 `json:"percentage_list"`
	Exercisetype         int       `json:"exercisetype"`
	ReptypeList          []int     `json:"reptype_list"`
	Name                 string    `json:"name"`
	ExerciseID           int       `json:"exercise_id"`
	Target               int       `json:"target"`
}

// PersonalTrainerApp returns the workout program in Personal Trainer App format
func PersonalTrainerApp(program programs.Program) []byte {

	workout := PTAWorkout{
		Explanation: program.Explanation,
		ShortName:   program.Name,
		Routinetype: "RC", // Where to store this workout in the app
		Category:    1,    // Not sure what this is in the app
		Name:        program.Name,
	}

	workout.fillDefaults()

	var dayNames []string // "Squat Day 1" etc
	var daysList [][]DaysList
	programSet := programs.NewProgramSet()

	// Programs are week based, so loop through all weeks to generate weights
	for week := 1; week <= program.Weeks; week++ {

		weekNames, weekWorkouts := program.Routine(week, &programSet) // ["Squat Day 1", "Bench Day 1"], [[session, session], [session, session]]

		for day := 0; day < len(weekNames); day++ {

			dayNames = append(dayNames, weekNames[day]) // "Squat Day 1"

			workoutDay := weekWorkouts[day] // [session, session]

			var sessionsList []DaysList
			for _, session := range workoutDay {

				// Set PTA fields based on the workouts
				var exerciseType, incrementType int

				switch session.IncrementType {
				case programs.IncrementWeightsProgramComplete:
					exerciseType = ExerciseTypeWendlerMainLift
					incrementType = IncrementTypeYes
				case programs.IncrementWeightsPerSession:
					exerciseType = ExerciseTypePercentage
					incrementType = IncrementTypeYes
				case programs.IncrementWeightsOff:
					exerciseType = ExerciseTypeWendlerMainLift
					incrementType = IncrementTypeNo
				}

				// Build items that require one per set
				reptypeList := []int{}
				weightksList := []float64{}
				weightslbList := []int{}
				repsList := []int{}
				percentageList := []float64{}

				lastSetAmrap := 0
				if session.LastSetIsAmrap {
					lastSetAmrap = 1
				}
				for set := 0; set < len(session.Sets.SetList); set++ {

					// Last set
					weightksList = append(weightksList, 2.5)
					weightslbList = append(weightslbList, 5)
					repsList = append(repsList, session.Sets.SetList[set].Reps)
					percentageList = append(percentageList, session.Sets.SetList[set].WeightPercentage)

					if set == len(session.Sets.SetList)-1 {
						reptypeList = append(reptypeList, lastSetAmrap) // Last set is AMRAP
					} else {
						reptypeList = append(reptypeList, 0)
					}

				}

				session := DaysList{
					Resttime1:       session.RestTimeSeconds,
					WeightskgList:   weightksList,
					RepsList:        repsList,
					LastsetsisAMRAP: session.LastSetIsAmrap,
					IncrementType:   incrementType,
					WeightslbList:   weightslbList,
					PercentageList:  percentageList,
					Exercisetype:    exerciseType,
					ReptypeList:     reptypeList,
					Name:            session.Lift.Name,
					ExerciseID:      session.Lift.ExerciseID,
					Target:          session.Lift.Target,
				}

				session.fillDefaults()

				sessionsList = append(sessionsList, session)
			}
			daysList = append(daysList, sessionsList)
		}
	}

	// Workout is complete, morph into usable JSON
	workout.DaysList = daysList
	workout.DayNames = dayNames
	workout.setDayNumber(len(dayNames))
	workout.setDaysPerWeek(len(dayNames) / program.Weeks)

	workoutJSON, err := json.MarshalIndent(workout, "", "  ")

	if err != nil {
		log.Fatalf(err.Error())
	}

	return workoutJSON
}

// FillDefaults sets default values for a workout
func (w *PTAWorkout) fillDefaults() {

	// setting default values
	// if no values present
	if w.ThemeColor == 0 {
		w.ThemeColor = 15138560
	}

	if w.Noofexercises == 0 {
		w.Noofexercises = 2 // No idea what this does
	}
}

// setDaysPerWeek sets values that should be workouts per week
func (w *PTAWorkout) setDaysPerWeek(daysPerWeek int) {
	w.Noofdays = daysPerWeek    // Set to days per week
	w.ProgramDays = daysPerWeek // Set to days per week
	w.Realdays = daysPerWeek    // Set to days per week
}

// SetDayNumber sets required values in workout to the number of days in a program
func (w *PTAWorkout) setDayNumber(days int) {
	w.Days = days
	w.WhenToUpdateWeights = days

	expanded := []bool{}

	for i := 0; i < days; i++ {
		expanded = append(expanded, true)
	}

	w.Expanded = expanded
}

// Fill in default values for Session
func (d *DaysList) fillDefaults() {

	d.SameWeight = true
	d.SameWeightWarmUp = true
	d.BarbellKg = 20
	d.BarbellLb = 45
	d.Incrementkg = 0
	d.ResttimeWarmup = 60
	d.SupersetNumber = -1
	d.AllsetsareAMRAP = false
	d.Type = 0

	if d.Failuresallowed == 0 {
		d.Failuresallowed = 1
	}

	if d.Currentweightlb == 0 {
		d.Currentweightlb = 2.5
		d.Currentweightkg = 2.5
	}

	if d.Resttime1 == 0 {
		d.Resttime1 = 90
	}

	d.Resttime2 = int(float32(d.Resttime1) * float32(2))
	d.Resttime3 = d.Resttime1 * 3

	if d.DeloadPercentage == 0 {
		d.DeloadPercentage = 10
	}

	if d.DependentIncrementID == 0 {
		d.DependentIncrementID = -1
		d.DependentExerciseID = -1
	}
}
