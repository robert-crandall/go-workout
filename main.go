package main

import (
	"main/outputs"
	"main/programs"
)

func main() {
	// Create base workout
	workoutPrograms := programs.GetPrograms()

	for _, program := range workoutPrograms {

		if program.Export {
			outputs.Markdown(program)
		}

	}
}
