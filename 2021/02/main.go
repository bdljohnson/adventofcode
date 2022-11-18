package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bdljohnson/advent-of-code/helpers"
)

func main() {
	input := helpers.Input.GetStringInput()
	inputInstructions := splitInstructions(input)
	runOne(inputInstructions)
	runTwo(inputInstructions)
}

type instruction struct {
	direction string
	length    int
}

func splitInstructions(input []string) []instruction {
	var inputInstructions []instruction
	for _, v := range input {
		row := strings.Split(v, " ")
		length, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal("Error converting input")
		}
		rowInstruction := instruction{
			direction: row[0],
			length:    length,
		}
		inputInstructions = append(inputInstructions, rowInstruction)
	}
	return inputInstructions
}

func runOne(input []instruction) {
	x := 0
	y := 0
	for _, v := range input {
		switch v.direction {
		case "up":
			y -= v.length
		case "down":
			y += v.length
		case "forward":
			x += v.length
		default:
			fmt.Println("invalid instruction")
			continue
		}
	}
	fmt.Printf("X: %v, Y: %v, VEC: %v\n", x, y, x*y)
}

func runTwo(input []instruction) {
	x := 0
	y := 0
	aim := 0

	for _, v := range input {
		switch v.direction {
		case "up":
			aim -= v.length
		case "down":
			aim += v.length
		case "forward":
			x += v.length
			y += v.length * aim
		}
	}
	fmt.Printf("X: %v, Y: %v, VEC: %v", x, y, x*y)
}
