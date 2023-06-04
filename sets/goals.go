package sets

type Goal int

const (
	Maintain Goal = iota
	Increase
	OneRM
	Lite
)

func (s Goal) String() string {
	switch s {
	case Maintain:
		return "Maintain"
	case Increase:
		return "Increase"
	case OneRM:
		return "OneRM"
	case Lite:
		return "Lite"
	}
	return "unknown"
}

func (s Goal) Adjustments() int {
	switch s {
	case Maintain:
		return 3
	case Increase:
		return 2
	case OneRM:
		return 2
	case Lite:
		return 4
	default:
		return 0
	}
}
