package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/exporters"
	"main/programs"
	"os"
	"time"
)

var (
	markdownExportEnabled        = true
	personalTrainerExportEnabled = true
)

func main() {
	// Cleanup current directory
	cleanUp()

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

func cleanUp() {
	files, err := ioutil.ReadDir("output/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err := os.Remove(fmt.Sprintf("output/%s", file.Name()))
		if err != nil {
			fmt.Println(err)
		}
	}
}
