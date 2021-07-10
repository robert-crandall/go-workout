package outputs

import (
	"fmt"
	"main/programs"
)

type result struct {
	output string
}

// Markdown returns the workout program in markdown format
func Markdown(program programs.Program) []byte {
	output := result{}

	output.addHeader(1, program.Name)
	output.addHeader(2, program.Explanation)

	for week := 0; week < program.Weeks; week++ {
		lastWeek := false
		if week == program.Weeks-1 {
			lastWeek = true
		}

		weekNames, weekWorkouts := program.Routine(week)

		for day := 0; day < program.DaysPerWeek; day++ {
			output.addText("") // blank line
			output.addText(boldText(weekNames[day]))

			workoutDay := weekWorkouts[day]
			for _, lift := range workoutDay {
				output.addText(fmt.Sprintf("* %s", lift.Lift.Name))

				for set := 0; set < len(lift.Set.RepsList); set++ {
					lastSet := false
					if set == len(lift.Set.RepsList)-1 {
						lastSet = true
					}
					var weightText, amrap, addWeight string
					if lift.Set.WeightsLBList != nil {
						weightText = fmt.Sprintf("%v lbs", lift.Set.WeightsLBList[set])
					} else {
						weightText = fmt.Sprintf("%v%%", lift.Set.PercentageList[set]*100)
					}
					if lift.Set.LastSetsIsAMRAP {
						amrap = "+"
					}

					if lift.IncrementType == programs.IncrementTypeYes && lastSet {
						if lift.ExerciseType == programs.ExerciseTypeWeightBased || lift.ExerciseType == programs.ExerciseTypePercentage {
							addWeight = "(Add weight)"
						} else if lastWeek {
							addWeight = "(Add weight)"
						}
					}

					output.addText(fmt.Sprintf("  * %d%s @ %s %s", lift.Set.RepsList[set], amrap, weightText, addWeight))
				}
			}
		}
	}

	fmt.Println(output.output)
	return []byte("Hello world")
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
