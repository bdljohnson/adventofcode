package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/bdljohnson/advent-of-code/helpers"
)

func main() {
	input := helpers.Input.GetStrInput()
	run_one(input)
}

func run_one(input []string) {
	var resultset []int
	temp := 0
	for _, v := range input {
		if len(v) > 0 {
			i, _ := strconv.Atoi(v)
			temp += i
		} else {
			resultset = append(resultset, temp)
			temp = 0
		}
	}
	sort.Ints(resultset)

	fmt.Println("Case 1: ", resultset[len(resultset)-1])
	fmt.Println("Case 2: ", resultset[len(resultset)-1]+resultset[len(resultset)-2]+resultset[len(resultset)-3])
}
