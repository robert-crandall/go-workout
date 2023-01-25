package sets

import (
	"math"
)

type Sets struct {
	SetList          []set
	Goal             Goal       // Goal for the workout
	LiftScheme       LiftScheme // Scheme for the workout. 3x5, 3x8, etc. Can override sets and reps with options.
	LastSetsIsAMRAP  bool       // Last set is As Many Reps as Possible // TODO - this should be private
	RestTimeSeconds  int        // Time between sets // TODO - this should be private
	setCount         int        // How many sets to perform
	repCount         int        // How many reps to perform, or starting number for (reverse) pyramid
	weightPercentage float64    // Percentage for weight, or starting percentage for (reverse) pyramid
}

type set struct {
	Reps             int
	WeightPercentage float64
}

type Options func(s *Sets)

func NewSets(goal Goal, liftScheme LiftScheme, opts ...Options) Sets {
	s := Sets{
		Goal:       goal,
		LiftScheme: liftScheme,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

func NewOptions() Options {
	return func(s *Sets) {}
}

func WithSetCount(setCount int) Options {
	return func(s *Sets) {
		s.setCount = setCount
	}
}

func WithRestTimer(restTimeSeconds int) Options {
	return func(s *Sets) {
		s.RestTimeSeconds = restTimeSeconds
	}
}

func WithRepCount(reps int) Options {
	return func(s *Sets) {
		s.repCount = reps
	}
}

func WithWeightPercentage(percentage float64) Options {
	return func(s *Sets) {
		// Round percentage to 2 decimal places
		s.weightPercentage = percentage
	}
}

func (s *Sets) setTargetWeight() {

	if s.setCount == 0 {
		s.setSetCount()
	}

	weightPercentage := percentageOf1RM(s.repCount)
	weightPercentage = weightPercentage * 0.97

	switch s.Goal {
	case Maintain:
		s.weightPercentage = weightPercentage - 0.05
	case Increase:
		s.weightPercentage = weightPercentage
	case Lite:
		s.weightPercentage = weightPercentage - 0.25
	default:
		s.weightPercentage = weightPercentage
	}
}

func (s *Sets) setRestTimer() {
	switch s.Goal {
	case Maintain:
		s.RestTimeSeconds = 90
	case Increase:
		s.RestTimeSeconds = 120
	case OneRM:
		s.RestTimeSeconds = 300
	case Lite:
		s.RestTimeSeconds = 90
	default:
		s.RestTimeSeconds = 90
	}
}

func (s *Sets) setSetCount() {
	s.setCount = s.LiftScheme.Sets()
}

func (s *Sets) setRepCount() {
	s.repCount = s.LiftScheme.Reps()
}

func (s Sets) GetProgram() Sets {

	if s.setCount == 0 {
		s.setSetCount()
	}

	if s.repCount == 0 {
		s.setRepCount()
	}

	if s.weightPercentage == 0 {
		s.setTargetWeight()
	}

	if s.RestTimeSeconds == 0 {
		s.setRestTimer()
	}

	s = s.Static()

	if s.LiftScheme.IsRPT() {
		s = s.RPT(2, 0.05)
	}

	if s.LiftScheme.Is1RM() {
		s = s.OneRepMaxTest()
	}

	return s
}

// Static configures the setList as it was defined by options
func (s Sets) Static(overrideSetCount ...int) Sets {
	var setList []set
	setCount := s.setCount

	if len(overrideSetCount) > 0 {
		setCount = overrideSetCount[0]
	}

	for i := 0; i < setCount; i++ {
		thisSet := set{
			Reps:             s.repCount,
			WeightPercentage: truncateNum(s.weightPercentage),
		}
		setList = append(setList, thisSet)
	}
	s.SetList = setList
	return s
}

// RPT configures the setList following a Reverse Pyramid Scheme, increasing and decreasing by the variables
func (s Sets) RPT(repIncrease int, decrementPercent float64) Sets {
	var setList []set

	for i := 0; i < s.setCount; i++ {
		thisSet := set{
			Reps:             s.repCount + (repIncrease * i),
			WeightPercentage: truncateNum(s.weightPercentage - (decrementPercent * float64(i))),
		}
		setList = append(setList, thisSet)
	}
	s.SetList = setList
	return s
}

// RPT configures the setList following a Reverse Pyramid Scheme, increasing and decreasing by the variables
func (s Sets) OneRepMaxTest() Sets {
	var setList []set

	for i := s.repCount; i > 0; i -= 2 {
		thisSet := set{
			Reps:             i,
			WeightPercentage: truncateNum(percentageOf1RM(i) * 0.97),
		}
		setList = append(setList, thisSet)
	}

	s.SetList = setList
	return s
}

// truncateNum is a helper function to round all floats to 2 decimal places
func truncateNum(num float64) (result float64) {
	return math.Round(num*100) / 100
}

// Returns a percentage of 1RM given the number of reps
func percentageOf1RM(reps int) float64 {
	// I realize this is something like (1 - 0.027*reps). I'm keeping it as a lookup table for easy understanding.
	switch reps {
	case 1:
		return 1
	case 2:
		return 0.97
	case 3:
		return 0.94
	case 4:
		return 0.91
	case 5:
		return 0.88
	case 6:
		return 0.85
	case 7:
		return 0.82
	case 8:
		return 0.79
	case 9:
		return 0.76
	case 10:
		return 0.73
	case 11:
		return 0.70
	case 12:
		return 0.67
	default:
		return 0.65
	}
}
