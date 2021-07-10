package sets

type Set struct {
	RepsList        []int
	WeightsLBList   []int     // For weight based lifts
	PercentageList  []float64 // For percentage based lifts
	LastSetsIsAMRAP bool      // Last set is As Many Reps as Possible
}

// FSL531 returns a 5x5 Set matching Wendler's 531 First Set Last program
// Weight is a percentage based on the week number
func FSL531(week int) Set {

	repsList := []int{}
	percentageList := []float64{}

	var percentage float64

	switch week {
	case 1:
		percentage = 0.70
	case 2:
		percentage = 0.725
	case 3:
		percentage = 0.75
	case 4:
		percentage = 0.60
	}

	reps := 5
	sets := 5

	for i := 0; i < sets; i++ {
		repsList = append(repsList, reps)
		percentageList = append(percentageList, percentage)
	}

	return Set{
		RepsList:        repsList,
		PercentageList:  truncateNumList(percentageList),
		LastSetsIsAMRAP: false,
	}
}

// Main531 returns a Set matching Wendler's 531 program
// Weight is a percentage based on week number
func Main531(week int) Set {

	repsList := []int{}
	percentageList := []float64{}
	lastSetIsAmrap := false

	switch week {
	case 1:
		repsList = []int{5, 5, 5}
		percentageList = []float64{0.65, 0.75, 0.85}
	case 2:
		repsList = []int{3, 3, 3}
		percentageList = []float64{0.70, 0.80, 0.90}
	case 3:
		repsList = []int{5, 3, 1}
		percentageList = []float64{0.75, 0.85, 0.95}
		lastSetIsAmrap = true
	case 4:
		repsList = []int{5, 5, 5}
		percentageList = []float64{0.40, 0.50, 0.60}
	}

	return Set{
		RepsList:        repsList,
		PercentageList:  truncateNumList(percentageList),
		LastSetsIsAMRAP: lastSetIsAmrap,
	}
}

// RPTIncreaseWeight returns a Reverse Pyramid Exercise, starting at 80%, with more weight for 3 weeks, then deload
// Lower is 4/6/8 reps
// Upper is 6/8/10 reps
// Increment applies after a cycle is complete
func RPTIncreaseWeight(week int, upper bool) Set {

	repsList := []int{}
	percentageList := []float64{}

	var base, rptExtraReps int
	var startingPercentage, weightDecrement, weeklyWeightIncrease float64

	// How much weight to add for next weeks
	weeklyWeightIncrease = 0.025

	if upper {
		base = 6
		rptExtraReps = 2
		startingPercentage = 0.80 // Based on 1RM tables for 8 reps
		weightDecrement = 0.05    // What percentage to decrease
	} else {
		base = 4
		rptExtraReps = 1
		startingPercentage = 0.85 // Based on 1RM tables for 6 reps
		weightDecrement = 0.05
	}

	// Last week - take it easy
	if week == 4 {
		startingPercentage = 0.65
		weeklyWeightIncrease = 0.0
	}

	// Each week, starting percentage is increased by weeklyWeightIncrease
	startingPercentage = startingPercentage + (weeklyWeightIncrease * float64(week-1))

	repsList = []int{base, base + rptExtraReps, base + (rptExtraReps * 2)}
	percentageList = []float64{startingPercentage, startingPercentage - weightDecrement, startingPercentage - (weightDecrement * 2)}

	return Set{
		RepsList:        repsList,
		PercentageList:  truncateNumList(percentageList),
		LastSetsIsAMRAP: false,
	}
}

// StaticSets returns an Xx5 exercise at a given percentage
// Week 4 is a deload week
func StaticSets(week, sets int, percentage float64) Set {

	repsList := []int{}
	percentageList := []float64{}

	// Last week - take it easy
	if week == 4 {
		percentage = 0.65
	}

	reps := 5

	for i := 0; i < sets; i++ {
		repsList = append(repsList, reps)
		percentageList = append(percentageList, percentage)
	}

	return Set{
		RepsList:        repsList,
		PercentageList:  truncateNumList(percentageList),
		LastSetsIsAMRAP: false,
	}
}

// StaticSetsIncreaseReps returns a sets x reps Exercise at 100%
// Reps are increased by increase for every set
func StaticSetsIncreaseReps(sets, reps, increase int) Set {

	repsList := []int{}
	percentageList := []float64{}

	for i := 0; i < sets; i++ {
		rep := reps + (i * increase)
		repsList = append(repsList, rep)
		percentageList = append(percentageList, 1.0)
	}

	return Set{
		RepsList:        repsList,
		PercentageList:  truncateNumList(percentageList),
		LastSetsIsAMRAP: false,
	}
}

// truncateNumList is a helper function to round all floats to 3 decimal places
func truncateNumList(numList []float64) (result []float64) {
	for _, v := range numList {
		result = append(result, float64(int(v*1000))/1000)
	}
	return
}
