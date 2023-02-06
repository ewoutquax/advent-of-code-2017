package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	universe := parseInput(input())

	fmt.Println("universe:", universe)
	fmt.Println("scanner at row-0:", universe.scanners[0])

	assert.Equal(t, 6, universe.maxRows)
	assert.Equal(t, 4, len(universe.scanners))
	assert.Equal(t, 3, universe.scanners[0].maxDepth)
	assert.Equal(t, 4, universe.scanners[6].maxDepth)
}

func TestMovePlayerThroughUniverse(t *testing.T) {
	universe := parseInput(input())
	universe.movePlayerToOtherSide()

	assert.Equal(t, 24, universe.Player.tripSeverity())
}

func input() []string {
	return []string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	}
}
