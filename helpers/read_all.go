package helpers

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

/*
	Read_All_Input reads all lines in a file and []string.
	Any further manipulation requires converting to the desired type.
*/
func Read_All_Input(filename string) []string {
	var input []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not open file:\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

/*
	Read_All_Input_Int will parse an input file and return an integer input set.
*/
func Read_All_Input_Int(filename string) []int {
	var input []int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not open file:\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Could not convert input:\n", err)
		}
		input = append(input, i)
	}

	return input
}
