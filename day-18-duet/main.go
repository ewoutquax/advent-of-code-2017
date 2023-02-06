package main

import (
	"fmt"

	"github.com/ewoutquax/advent-of-code-2017/day-18/puzzle"
	"github.com/ewoutquax/advent-of-code-2022/utils"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1())
	fmt.Println("Result of part-2: ", solvePart2())
}

func solvePart1() int {
	return puzzle.FindLastPlayedValue(utils.ReadFileAsLines())
}

func solvePart2() int {
	fmt.Println("Placeholder for part 2")
	return 0
}
