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
