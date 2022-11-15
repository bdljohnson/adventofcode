package helpers

import (
	"bufio"
	"flag"
	"log"
)

// InputData represents various ways of accessing advent of code input data
type InputData struct {
	integerInput  []int //
	stringInput   []string
	inputStreamer bufio.Scanner
}

func (i *InputData) GetIntInput() []int {
	return i.integerInput
}

var Input = InputData{}

func init() {
	var filename string
	var inputType string
	flag.StringVar(&filename, "filename", "", "Filename of input")
	flag.StringVar(&inputType, "input-type", "integer", "Input type")
	flag.Parse()

	switch inputType {
	case "integer":
		Input.integerInput = processInt(filename)
	default:
		log.Fatal("Invalid input type")
	}

}

func processInt(filename string) []int {
	input := Read_All_Input_Int(filename)
	return input
}
