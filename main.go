package main

import (
	"fmt"
	"io/ioutil"
	"main/exporters"
	"main/programs"
	"time"
)

var (
	markdownExportEnabled        = true
	personalTrainerExportEnabled = true
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

			if personalTrainerExportEnabled {
				personalTrainerJSON := exporters.PersonalTrainerApp(program)
				outputName := fmt.Sprintf("output/%s-%d.workout", program.Name, time.Now().UnixNano())
				_ = ioutil.WriteFile(outputName, personalTrainerJSON, 0644)
			}
		}

	}
}
