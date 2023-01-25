package sets

type LiftScheme int

const (
	FiveByFive LiftScheme = iota
	ThreeByFive
	ThreeByEight
	RptFourSets
	RptFiveSets
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
		return "RPT 4 sets"
	case RptFiveSets:
		return "RPT 5 sets"
	}
	return "unknown"
}
