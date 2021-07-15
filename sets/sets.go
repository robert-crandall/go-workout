package sets

import (
	"math"
)

type Sets struct {
	SetList          []set
	LastSetsIsAMRAP  bool    // Last set is As Many Reps as Possible
	RestTimeSeconds  int     // Time between sets
	setCount         int     // How many sets to perform
	repCount         int     // How many reps to perform, or starting number for (reverse) pyramid
	weightPercentage float64 // Percentage for weight, or starting percentage for (reverse) pyramid
}

type set struct {
	Reps             int
	WeightPercentage float64
}

type Options func(s *Sets)

func NewSets(opts ...Options) Sets {
	s := Sets{
		setCount:        3,
		RestTimeSeconds: 90,
		repCount:        5,
	}
	for _, opt := range opts {
		opt(&s)
	}
	return s
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
		s.weightPercentage = percentage
	}
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

// truncateNum is a helper function to round all floats to 3 decimal places
func truncateNum(num float64) (result float64) {
	return math.Round(num*1000) / 1000
}
