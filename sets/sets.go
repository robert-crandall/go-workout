package sets

import (
	"math"
)

type Sets struct {
	SetList          []set
	Goal             Goal       // Goal for the workout
	LiftScheme       LiftScheme // Scheme for the workout. 3x5, 3x8, etc. Can override sets and reps with options.
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

func (s *Sets) goalWeight(reps int) float64 {
	if s.weightPercentage > 0 {
		return s.weightPercentage
	}

	switch s.Goal {
	case Maintain:
		return percentageOf1RM(reps + 2)
	case Increase:
		return percentageOf1RM(reps + 1)
	case OneRM:
		return percentageOf1RM(reps + 1)
	case Lite:
		return percentageOf1RM(reps + 4)
	default:
		return percentageOf1RM(reps)
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
			WeightPercentage: s.goalWeight(s.repCount),
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
			WeightPercentage: s.goalWeight(s.repCount + (repIncrease * i)),
		}
		setList = append(setList, thisSet)
	}
	s.SetList = setList
	return s
}

// Aim to to 1 set at 100% of 1RM
func (s Sets) OneRepMaxTest() Sets {
	var setList []set

	for i := s.repCount; i > 0; i -= 2 {
		thisSet := set{
			Reps:             i,
			WeightPercentage: s.goalWeight(i),
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
	// Source of table is https://strengthlevel.com/one-rep-max-calculator/
	// It seemed to change last time I compared values. Either that or I grabbed from a different source initially.
	switch reps {
	case 1:
		return 1
	case 2:
		return 0.97
	case 3:
		return 0.94
	case 4:
		return 0.92
	case 5:
		return 0.89
	case 6:
		return 0.86
	case 7:
		return 0.83
	case 8:
		return 0.81
	case 9:
		return 0.78
	case 10:
		return 0.75
	case 11:
		return 0.73
	case 12:
		return 0.71
	case 13:
		return 0.70
	case 14:
		return 0.68
	case 15:
		return 0.67
	default:
		return 0.65
	}
}
