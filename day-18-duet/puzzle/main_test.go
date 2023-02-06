package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	program := parseInput(input(), v1)

	assert.Equal(t, 10, len(program.statements))
	assert.Equal(t, 7, len(program.values))

	assert.Equal(t, 1, program.values["1"].getValue())
	assert.Equal(t, -2, program.values["-2"].getValue())
}

func TestUpdateStateProgram(t *testing.T)      {}
func TestDetectEndStateRecovered(t *testing.T) {}
func TestDetectEndStateDeadlock(t *testing.T)  {}

func TestRunningTwoPrograms(t *testing.T) {
	occurences := FindOccurencesProgram1Sends(inputPart2())

	assert.Equal(t, 1, occurences)
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

func inputPart2() []string {
	return []string{
		"snd 1",
		"snd 2",
		"snd p",
		"rcv a",
		"rcv b",
		"rcv c",
		"rcv d",
	}
}
