package programs

type Program struct {
	Noofdays    int // Workouts per week
	Explanation string
	ShortName   string
	RoutineType string // Category for organization
	Category    int    // No idea
	Name        string
	Weeks       int  // Weeks until a complete cycle
	Export      bool // Whether to export this program or not
}

var (
	programs []Program
)

func GetPrograms() []Program {
	return programs
}
