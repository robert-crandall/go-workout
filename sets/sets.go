package sets

type Set struct {
	repsList        []int
	weightsLBList   []int     // For weight based lifts
	percentageList  []float64 // For percentage based lifts
	lastSetsIsAMRAP bool      // Last set is As Many Reps as Possible
}

// FSL531 returns a 5x5 Set matching Wendler's 531 First Set Last program
// Weight is a percentage based on the week number
func FSL531(week int) Set {

	repsList := []int{}
	percentageList := []float64{}

	var percent float64

	switch week {
	case 1:
		percent = 0.70
	case 2:
		percent = 0.725
	case 3:
		percent = 0.75
	case 4:
		percent = 0.60
	}

	reps := 5
	sets := 5

	for i := 0; i < sets; i++ {
		repsList = append(repsList, reps)
		percentageList = append(percentageList, percent)
	}

	return Set{
		repsList:        repsList,
		percentageList:  percentageList,
		lastSetsIsAMRAP: false,
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
		repsList:        repsList,
		percentageList:  percentageList,
		lastSetsIsAMRAP: lastSetIsAmrap,
	}
}
