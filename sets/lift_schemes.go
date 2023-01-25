package sets

import (
	"fmt"
	"os"
)

type LiftScheme int

const (
	FiveByFive LiftScheme = iota
	ThreeByFive
	ThreeByEight
	RptFourSets
	RptFiveSets
	OneRepMaxTest
)

func (s LiftScheme) String() string {
	switch s {
	case FiveByFive:
		return "5x5"
	case ThreeByFive:
		return "3x5"
	case ThreeByEight:
		return "3x8"
	case RptFourSets:
		return "RPT"
	case RptFiveSets:
		return "RPT"
	case OneRepMaxTest:
		return "1RM Test"
	}
	return "unknown"
}

func (s LiftScheme) IsRPT() bool {
	switch s {
	case RptFourSets, RptFiveSets:
		return true
	}
	return false
}

func (s LiftScheme) Is1RM() bool {
	switch s {
	case OneRepMaxTest:
		return true
	}
	return false
}

func (s LiftScheme) Sets() int {
	switch s {
	case FiveByFive, RptFiveSets:
		return 5
	case ThreeByFive, ThreeByEight, OneRepMaxTest:
		return 3
	case RptFourSets:
		return 4
	}
	fmt.Printf("%s - Lift Scheme does not have a Set Count\n", s.String())
	os.Exit(1)
	return 0
}

func (s LiftScheme) Reps() int {
	switch s {
	case FiveByFive, ThreeByFive, OneRepMaxTest:
		return 5
	case ThreeByEight:
		return 8
	}
	fmt.Printf("%s - Lift Scheme does not have a Rep Count\n", s.String())
	os.Exit(1)
	return 0
}
