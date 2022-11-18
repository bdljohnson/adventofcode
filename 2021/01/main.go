package main

import (
	"fmt"

	"github.com/bdljohnson/advent-of-code/helpers"
)

func run_two(input []int) {
	result := 0
	for i, v := range input {
		if i < 3 {
			continue
		}
		window1 := input[i-3] + input[i-2] + input[i-1]
		window2 := input[i-2] + input[i-1] + v
		if window2 > window1 {
			result += 1
		}
	}
	fmt.Printf("Input length: %v, Total Increases: %v\n", len(input), result)
}

func run_one(input []int) {
	result := 0
	for i := range input {
		if i > 0 && input[i] > input[i-1] {
			result += 1
		}
	}
	fmt.Printf("Total increases for run one: %v\n", result)
}

func main() {
	input := helpers.Input.GetIntInput()
	run_one(input)
	run_two(input)
}
