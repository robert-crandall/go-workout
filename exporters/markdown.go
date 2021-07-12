package exporters

import (
	"fmt"
	"main/programs"
	"math"
)

type result struct {
	output string
}

// Markdown returns the workout program in markdown format
func Markdown(program programs.Program) []byte {
	output := result{}

	output.addHeader(1, program.Name)
	output.addHeader(2, program.Explanation)

	for week := 1; week <= program.Weeks; week++ {
		lastWeek := false
		if week == program.Weeks {
			lastWeek = true
		}

		output.addText("")
		output.addHeader(3, fmt.Sprintf("Week %d", week))
		weekNames, weekWorkouts := program.Routine(week)

		for day := 0; day < program.DaysPerWeek; day++ {
			output.addText("") // blank line
			output.addText(boldText(weekNames[day]))

			workoutDay := weekWorkouts[day]
			for _, lift := range workoutDay {
				output.addText(fmt.Sprintf("* %s", lift.Lift.Name))

				for set := 0; set < len(lift.Session.RepsList); set++ {
					lastSet := false
					if set == len(lift.Session.RepsList)-1 {
						lastSet = true
					}
					var weightText, amrap, addWeight string
					if lift.Session.WeightsLBList != nil {
						weightText = fmt.Sprintf("%v lbs", lift.Session.WeightsLBList[set])
					} else {
						weightText = fmt.Sprintf("%v%%", truncateNum(lift.Session.PercentageList[set]*100))
					}
					if lift.Session.LastSetsIsAMRAP {
						amrap = "+"
					}

					if lastSet {
						if lift.IncrementType == programs.IncrementWeightsPerSession {
							addWeight = "(Add weight)"
						} else if lastWeek && lift.IncrementType == programs.IncrementWeightsProgramComplete {
							addWeight = "(Add weight)"
						}
					}

					output.addText(fmt.Sprintf("  * %d%s @ %s %s", lift.Session.RepsList[set], amrap, weightText, addWeight))
				}
			}
		}
	}

	return []byte(output.output)
}

func (r *result) addHeader(level int, text string) {
	var prefix string
	for i := 0; i < level; i++ {
		prefix = prefix + "#"
	}
	r.output = r.output + fmt.Sprintf("%s %s\n", prefix, text)
}

func (r *result) addText(text string) {
	r.output = r.output + fmt.Sprintf("%s\n", text)
}

func (r *result) addRow(rowText []string) {
	for i := 0; i < len(rowText); i++ {
		r.output = r.output + fmt.Sprintf("| %s ", rowText[i])
	}
	r.output = r.output + fmt.Sprintf(" |\n")
}

func boldText(text string) string {
	return fmt.Sprintf("**%s**", text)
}

// truncateNumList is a helper function to round all floats to 1 decimal places
func truncateNum(num float64) float64 {
	return math.Round(num*10) / 10
}
