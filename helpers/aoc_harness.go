package helpers

import (
	"bufio"
	"flag"
	"log"
)

var _testing = false

// InputData represents various ways of accessing advent of code input data
type InputData struct {
	integerInput  []int //
	stringInput   []string
	inputStreamer bufio.Scanner
}

func (i *InputData) GetIntInput() []int {
	return i.integerInput
}

func (i *InputData) GetStrInput() []string {
	return i.stringInput
}

var Input = InputData{}

func init() {
	if _testing {
		return
	}
	var filename string
	var inputType string
	flag.StringVar(&filename, "filename", "", "Filename of input")
	flag.StringVar(&inputType, "input-type", "integer", "Input type")
	flag.Parse()

	if len(filename) == 0 {
		log.Fatal("Requires a filename!")
	}

	switch inputType {
	case "integer":
		Input.integerInput = processInt(filename)
	case "string":
		Input.stringInput = processStr(filename)
	default:
		log.Fatal("Invalid input type")
	}

}

func processInt(filename string) []int {
	input := ReadAllInputInt(filename)
	return input
}

func processStr(filename string) []string {
	input := ReadAllInput(filename)
	return input
}
