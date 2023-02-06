package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecStatements(t *testing.T) {
	program := parseInput(input(), v1)

	program.execStatements()

	assert.Equal(t, 4, program.valueLastPlayed)
}

func TestWillSendStateUpdateReceive(t *testing.T) {
	c1 := make(chan struct{ stateProgram }, 5)
	c2 := make(chan int)
	program := parseInput([]string{"rcv a"}, v2)

	program.chanToUniverse = c1
	program.chanToOtherProgram = c2

	fmt.Println("c1:", c1)

	program.statements[0].exec(&program)

	fmt.Println("update:", <-c1)

	assert.True(t, false)
}

func TestWillSendStateUpdateRecovered(t *testing.T) {}
