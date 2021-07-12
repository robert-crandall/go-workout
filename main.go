package main

import (
	"fmt"
	"io/ioutil"
	"main/exporters"
	"main/programs"
)

var (
	markdownExportEnabled = true
)

func main() {
	// Create base workout
	workoutPrograms := programs.GetPrograms()

	for _, program := range workoutPrograms {

		if program.Export {
			if markdownExportEnabled {
				markdownBytes := exporters.Markdown(program)
				outputName := fmt.Sprintf("output/%s.md", program.Name)
				_ = ioutil.WriteFile(outputName, markdownBytes, 0644)
			}
		}

	}
}
