package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveScanner(t *testing.T) {
	universe := parseInput([]string{"0: 2"})

	scanner0 := universe.scanners[0]

	scanner0.move()
	assert.Equal(t, 1, scanner0.currentLayer)
	assert.Equal(t, ScannerDirectionDown, scanner0.Direction)

	scanner0.move()
	assert.Equal(t, 0, scanner0.currentLayer)
	assert.Equal(t, ScannerDirectionUp, scanner0.Direction)
}

func TestWillHitPlayerForDelay(t *testing.T) {
	var scanner Scanner

	scanner = newScanner(1, 2)

	assert.False(t, scanner.willCatchPlayerForDelay(0))
	assert.True(t, scanner.willCatchPlayerForDelay(1))
	assert.False(t, scanner.willCatchPlayerForDelay(2))
	assert.True(t, scanner.willCatchPlayerForDelay(3))

	scanner = newScanner(5, 3)

	assert.False(t, scanner.willCatchPlayerForDelay(0))
	assert.True(t, scanner.willCatchPlayerForDelay(1))
	assert.False(t, scanner.willCatchPlayerForDelay(2))
	assert.False(t, scanner.willCatchPlayerForDelay(3))
	assert.True(t, scanner.willCatchPlayerForDelay(4))
}
