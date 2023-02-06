package puzzle

import (
	"strings"

	"github.com/ewoutquax/advent-of-code-2022/utils"
)

func CalculateTripSeverity(lines []string) int {
	universe := parseInput(lines)

	universe.movePlayerToOtherSide()

	return universe.Player.tripSeverity()
}

func FindLowestDelayForUncaught(lines []string) int {
	universe := parseInput(lines)

	return universe.findLowestDelayForUncaught()
}

func parseInput(lines []string) (universe Universe) {
	var parts []string
	universe.initUniverse(len(lines))

	for _, line := range lines {
		parts = strings.Split(line, ": ")
		scanner := newScanner(
			utils.ConvStrToI(parts[0]),
			utils.ConvStrToI(parts[1]),
		)
		universe.maxRows = scanner.row
		universe.scanners[scanner.row] = &scanner
	}

	return
}
