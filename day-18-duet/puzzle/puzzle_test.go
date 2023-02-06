package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	universe := parseInput(input())

	assert.Equal(t, 10, len(universe.statements))
	assert.Equal(t, 7, len(universe.values))

	assert.Equal(t, 1, universe.values["1"].getValue())
	assert.Equal(t, -2, universe.values["-2"].getValue())
}

func TestExecStatements(t *testing.T) {
	universe := parseInput(input())

	universe.execStatements()

	assert.Equal(t, 4, universe.valueLastPlayed)
}

func input() []string {
	return []string{
		"set a 1",
		"add a 2",
		"mul a a",
		"mod a 5",
		"snd a",
		"set a 0",
		"rcv a",
		"jgz a -1",
		"set a 1",
		"jgz a -2",
	}
}
